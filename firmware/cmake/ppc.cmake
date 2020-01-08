set(ppc ${CONAN_PPC_ROOT})
include(MQTT)

file(GLOB SOURCE_FILES ${ppc}/src/*)

add_library(ppc STATIC ${SOURCE_FILES})
target_include_directories(ppc PRIVATE ${ppc}/include ${MQTT}/include ${PLATFORM_CXX_INCLUDES})
target_compile_options(ppc PRIVATE "$<$<CONFIG:ALL>:${PLATFORM_CXX_FLAGS}>")
target_compile_definitions(ppc PRIVATE ${PLATFORM_CXX_DEFS})
add_dependencies(ppc MQTT)
