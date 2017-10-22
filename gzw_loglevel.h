#ifndef GOOPENZWAVE_LOGLEVEL
#define GOOPENZWAVE_LOGLEVEL

#include <stdint.h>
#include <stdbool.h>
#include <stddef.h>

#ifdef __cplusplus
#include <platform/Log.h>
extern "C" {
#endif

// enum loglevel
typedef enum {
    loglevel_invalid = 0,
    loglevel_none,
    loglevel_always,
    loglevel_fatal,
    loglevel_error,
    loglevel_warning,
    loglevel_alert,
    loglevel_info,
    loglevel_detail,
    loglevel_debug,
    loglevel_streamdetail,
    loglevel_internal
} loglevel_t;

#ifdef __cplusplus
}
OpenZWave::LogLevel loglevel_toLogLevel(loglevel_t level);
#endif

#endif // define GOOPENZWAVE_LOGLEVEL
