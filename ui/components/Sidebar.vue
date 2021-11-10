<template>
  <div class="min-h-screen bg-gray-900 max-w-xs flex flex-col justify-between"
       :class="{ 'w-1/2': !collapsed, 'min-w-xs': !collapsed }">
    <div>
      <div class="bg-gray-900 px-6 py-8 text-2xl text-gray-100 font-bold flex items-center">
        <svg viewBox="0 0 300 300"
             version="1.1"
             xmlns="http://www.w3.org/2000/svg"
             xmlns:xlink="http://www.w3.org/1999/xlink"
             xml:space="preserve"
             xmlns:serif="http://www.serif.com/"
             class="fill-current w-12 h-12"
             style="fill-rule:evenodd;clip-rule:evenodd;stroke-linejoin:round;stroke-miterlimit:2;">
          <g id="Layer1">
            <path d="M219.986,239.429C200.751,255.275 176.107,264.797 149.257,264.797C87.788,264.797 37.882,214.891 37.882,153.422C37.882,91.952 87.788,42.046 149.257,42.046C210.727,42.046 260.633,91.952 260.633,153.422C260.633,178.194 252.528,201.088 238.819,219.585L224.545,205.286C234.741,190.549 240.705,172.676 240.705,153.422C240.705,102.95 199.729,61.974 149.257,61.974C98.786,61.974 57.81,102.95 57.81,153.422C57.81,203.893 98.786,244.869 149.257,244.869C170.607,244.869 190.257,237.537 205.814,225.232L219.986,239.429Z" />
            <g transform="matrix(-1,-0.000749452,0.000749452,-1,264.853,303.695)">
              <path d="M151.242,180.382L114.988,180.63L133.115,126.574L169.369,126.574L151.242,180.382Z" />
            </g>
            <g transform="matrix(-1,-0.000749452,0.000749452,-1,319.274,303.695)">
              <path d="M151.242,180.382L114.988,180.63L133.115,126.574L169.369,126.574L151.242,180.382Z" />
            </g>
          </g>
        </svg>
        <div class="ml-2"
              v-if="!collapsed">
          VirtM
          <div class="text-xs text-gray-400">
            {{ version }}
          </div>
        </div>
      </div>
      <nav class="py-0 px-4">
        <sidebar-item name="Home"
                      :collapsed="collapsed" to="/">
          <svg xmlns="http://www.w3.org/2000/svg"
               class="h-6 w-6"
               fill="none"
               viewBox="0 0 24 24"
               stroke="currentColor">
            <path stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
          </svg>
        </sidebar-item>
        <divider />
        <sidebar-section name="Virtual machines"
                         shortname="VMs"
                         :collapsed="collapsed">
          <sidebar-item name="Machines"
                        :collapsed="collapsed"
                        to="/machines">
            <svg xmlns="http://www.w3.org/2000/svg"
                 class="h-6 w-6"
                 fill="none"
                 viewBox="0 0 24 24"
                 stroke="currentColor">
              <path stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M5 12h14M5 12a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v4a2 2 0 01-2 2M5 12a2 2 0 00-2 2v4a2 2 0 002 2h14a2 2 0 002-2v-4a2 2 0 00-2-2m-2-4h.01M17 16h.01" />
            </svg>
          </sidebar-item>
          <sidebar-item name="SSH Keys"
                        :collapsed="collapsed" to="/ssh-keys">
            <svg xmlns="http://www.w3.org/2000/svg"
                 class="h-6 w-6"
                 fill="none"
                 viewBox="0 0 24 24"
                 stroke="currentColor">
              <path stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M15 7a2 2 0 012 2m4 0a6 6 0 01-7.743 5.743L11 17H9v2H7v2H4a1 1 0 01-1-1v-2.586a1 1 0 01.293-.707l5.964-5.964A6 6 0 1121 9z" />
            </svg>
          </sidebar-item>
          <sidebar-item name="Images"
                        :collapsed="collapsed" to="/images">
            <svg xmlns="http://www.w3.org/2000/svg"
                 class="h-6 w-6"
                 fill="none"
                 viewBox="0 0 24 24"
                 stroke="currentColor">
              <path stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M4 5a1 1 0 011-1h14a1 1 0 011 1v2a1 1 0 01-1 1H5a1 1 0 01-1-1V5zM4 13a1 1 0 011-1h6a1 1 0 011 1v6a1 1 0 01-1 1H5a1 1 0 01-1-1v-6zM16 13a1 1 0 011-1h2a1 1 0 011 1v6a1 1 0 01-1 1h-2a1 1 0 01-1-1v-6z" />
            </svg>
          </sidebar-item>
        </sidebar-section>
      </nav>
    </div>
    <div class="p-4">
      <button class="bg-gray-900 rounded-lg hover:text-gray-100 hover:bg-gray-800 text-gray-400 px-5 py-3 inline-block" @click="collapsed = !collapsed">
        <template v-if="collapsed">
          <svg xmlns="http://www.w3.org/2000/svg"
               class="h-6 w-6"
               fill="none"
               viewBox="0 0 24 24"
               stroke="currentColor">
            <path stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M13 5l7 7-7 7M5 5l7 7-7 7" />
          </svg>
        </template>
        <template v-else>
          <svg xmlns="http://www.w3.org/2000/svg"
               class="h-6 w-6"
               fill="none"
               viewBox="0 0 24 24"
               stroke="currentColor">
            <path stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M11 19l-7-7 7-7m8 14l-7-7 7-7" />
          </svg>
        </template>
      </button>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    version() {
      return this.$store.state.api.version
    },
  },
  data() {
    return {
      collapsed: window.innerWidth < 1024,
    };
  },
  mounted() {
    this.$store.dispatch('api/connect')
  }
};
</script>