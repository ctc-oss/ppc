#include <ppc.h>
#include <Particle.h>

static const String mqttcloudid = spark_deviceID().remove(0, 19);
static const std::string Fn(String::format("/F/%s/", mqttcloudid.c_str()));
static const std::string E("/E/");

using std::placeholders::_1;
using std::placeholders::_2;
using std::placeholders::_3;

namespace ppc
{
   MQTTCloud::MQTTCloud(const char* broker, uint16_t port)
         : client(const_cast<char*>(broker), port, 5, std::bind(&MQTTCloud::callback, this, _1, _2, _3)) {
   }

   MQTTCloud::~MQTTCloud() {
      if(client.isConnected()) client.disconnect();
   }

   bool MQTTCloud::publish(const char* eventName, const char* data, PublishFlags flags) {
      const auto e = E + eventName;
      return client.publish(e.c_str(), data);
   }

   bool MQTTCloud::subscribe(const char* eventName, EventHandler handler, Spark_Subscription_Scope_TypeDef scope) {
      const auto e = E + eventName;
      handlers[e] = std::make_shared<CloudEventHandler>(handler);
      return client.subscribe(e.c_str());
   }

   bool MQTTCloud::function(const char* funcKey, user_function_int_str_t* func) {
      return function(funcKey, std::function<user_function_int_str_t>(func));
   }

   bool MQTTCloud::function(const char* funcKey, const user_std_function_int_str_t& func, void* reserved) {
      const auto e = Fn + funcKey;
      functions[e] = std::make_shared<CloudFunc>(func);
      return client.subscribe(e.c_str());
   }

   void MQTTCloud::callback(char* t, uint8_t* d, uint32_t c) {
      const std::string e(t);
      const std::string p(reinterpret_cast< char const* >(d), c);

      if(e.rfind(Fn, 0) == 0) {     // Function
         const auto v = functions.find(e);
         if(v != functions.end()) {
            (*v->second)(p.c_str());
         }
      }
      else if(e.rfind(E, 0) == 0) { // Event
         const auto v = handlers.find(e);
         if(v != handlers.end()) {
            (*v->second)(e.substr(E.length()).c_str(), p.c_str());
         }
      }
   }

   bool MQTTCloud::connect(const char* id) {
      if(!id) return client.connect(mqttcloudid);
      else return client.connect(id);
   }
}
