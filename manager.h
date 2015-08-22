#ifndef GOZWAVE_MANAGER
#define GOZWAVE_MANAGER

#ifdef __cplusplus
extern "C" {
#endif

	// Types.
	typedef void* manager_t;

	// Static public member functions.
	manager_t manager_create();
	manager_t manager_get();
	void manager_destroy();
	const char* manager_getVersionAsString();
	const char* manager_getVersionLongAsString();

#ifdef __cplusplus
}
#endif

#endif // define GOZWAVE_MANAGER
