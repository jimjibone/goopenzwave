#include "manager.h"
#include <iostream>
#include <openzwave/Manager.h>

manager_t manager_init() {
	int* ret = new int();
	return (void*)ret;
}

void manager_free(manager_t f) {
	int* foo = (int*)f;
	delete foo;
}

void manager_bar(manager_t f) {
	int* foo = (int*)f;
	*foo = *foo + 1;

    std::cout << "manager printing numbers: " << OpenZWave::Manager::getVersionAsString() << std::endl;
}
