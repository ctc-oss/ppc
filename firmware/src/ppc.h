#pragma once

#include <Particle.h>
#include <MQTT.h>
#include <map>

typedef user_std_function_int_str_t CloudFunc;
typedef std::shared_ptr<CloudFunc> CloudFuncPtr;

typedef std::function<void(const char*, const char*)> CloudEventHandler;
typedef std::shared_ptr<CloudEventHandler> CloudEventHandlerPtr;

namespace ppc
{
   struct MQTTCloud
   {
      explicit MQTTCloud(const char* broker, uint16_t port = 1883);
      ~MQTTCloud();

      bool publish(const char* eventName, const char* data, PublishFlags flags);
      bool subscribe(const char* eventName, EventHandler handler, Spark_Subscription_Scope_TypeDef scope = MY_DEVICES);

      template<typename T, class ... Types>
      bool function(const T& name, Types ... args);
      bool function(const char* funcKey, user_function_int_str_t func, void* reserved = nullptr);
      bool function(const char* funcKey, user_std_function_int_str_t& func, void* reserved = nullptr);

      bool loop() { return client.loop(); }
      bool connect(const char* id) { return client.connect(id); }
      bool isConnected() { return client.isConnected(); }

   private:
      MQTT client;
      std::map<std::string, CloudFuncPtr> functions;
      std::map<std::string, CloudEventHandlerPtr> handlers;

      void callback(char* t, uint8_t* d, uint32_t c);
   };
}
