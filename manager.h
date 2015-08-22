#ifndef GOZWAVE_MANAGER
#define GOZWAVE_MANAGER

#ifdef __cplusplus
extern "C" {
#endif

	typedef void* manager_t;
	manager_t manager_init(void);
	void manager_free(manager_t);
	void manager_bar(manager_t);

#ifdef __cplusplus
}
#endif

#endif // define GOZWAVE_MANAGER
