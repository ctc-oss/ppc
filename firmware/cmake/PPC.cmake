set(PPC ${CONAN_PPC_ROOT})
include(MQTT)

file(GLOB SOURCE_FILES ${PPC}/src/*)

add_library(PPC STATIC ${SOURCE_FILES})
target_include_directories(PPC PRIVATE ${PPC}/include ${MQTT}/include ${PLATFORM_CXX_INCLUDES})
target_compile_options(PPC PRIVATE "$<$<CONFIG:ALL>:${PLATFORM_CXX_FLAGS}>")
target_compile_definitions(PPC PRIVATE ${PLATFORM_CXX_DEFS})
add_dependencies(PPC MQTT)
