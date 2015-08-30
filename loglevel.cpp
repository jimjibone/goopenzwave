#include "loglevel.h"

OpenZWave::LogLevel loglevel_toLogLevel(loglevel_t level)
{
    OpenZWave::LogLevel loglevel;
    switch (level) {
    case loglevel_invalid:
        loglevel = OpenZWave::LogLevel_Invalid;
        break;
    case loglevel_none:
        loglevel = OpenZWave::LogLevel_None;
        break;
    case loglevel_always:
        loglevel = OpenZWave::LogLevel_Always;
        break;
    case loglevel_fatal:
        loglevel = OpenZWave::LogLevel_Fatal;
        break;
    case loglevel_error:
        loglevel = OpenZWave::LogLevel_Error;
        break;
    case loglevel_warning:
        loglevel = OpenZWave::LogLevel_Warning;
        break;
    case loglevel_alert:
        loglevel = OpenZWave::LogLevel_Alert;
        break;
    case loglevel_info:
        loglevel = OpenZWave::LogLevel_Info;
        break;
    case loglevel_detail:
        loglevel = OpenZWave::LogLevel_Detail;
        break;
    case loglevel_debug:
        loglevel = OpenZWave::LogLevel_Debug;
        break;
    case loglevel_streamdetail:
        loglevel = OpenZWave::LogLevel_StreamDetail;
        break;
    case loglevel_internal:
        loglevel = OpenZWave::LogLevel_Internal;
        break;
    default:
        loglevel = OpenZWave::LogLevel_Invalid;
        break;
    }
    return loglevel;
}
