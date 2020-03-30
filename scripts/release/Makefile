OS_ARCH = $(findstring $(OS_TYPE), $(OPERATING_SYSTEMS))-$(findstring $(ARCH), $(OS_ARCHITECTURES))

#package-$(OS_TYPE)-$(ARCH):
package-$(OS_ARCH):
	echo TODO - Package

#setup-$(OS_TYPE)-$(ARCH):
setup-$(OS_ARCH):
	@echo Setup
	./common/setup.sh rel/beta

#package: package-$(findstring $(OS_TYPE), $(OPERATING_SYSTEMS))-$(findstring $(ARCH), $(OS_ARCHITECTURES))

#setup: setup-$(findstring $(OS_TYPE), $(OPERATING_SYSTEMS))-$(findstring $(ARCH), $(OS_ARCHITECTURES))

package: package-$(OS_ARCH)
setup: setup-$(OS_ARCH)

