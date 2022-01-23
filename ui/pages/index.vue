<template>
  <div class="max-w-2xl">
    <div class="h-10 flex justify-between items-center">
      <div class="flex items-center text-gray-400">
        <svg viewBox="0 0 10 11"
             class="h-3 mr-4 transform rotate-90"
             style="fill: currentColor">
          <g transform="matrix(-0.000438321,-0.925173,1,-0.000473772,-155.809,482.889)">
            <path d="M516.305,156.035L521.86,165.665L510.749,165.665L516.305,156.035Z" />
          </g>
        </svg>
        <h1 class="font-mono uppercase">
          Home
        </h1>
      </div>
    </div>
    <div class="mt-4">
      <div class="overflow-x-hidden flex flex-col items-end">
        <svg width="728"
             height="112"
             class="right-0">
          <g v-for="(week, index) in activitySummary"
             :key="index"
             :transform="`translate(${14 * index}, 0)`">
            <rect v-for="(score, index) in week"
                  :key="index"
                  width="10"
                  height="10"
                  x="0"
                  :y="13 * index"
                  rx="2"
                  ry="2"
                  class="fill-current"
                  :class="[ score > 0 ? 'text-oxide-400' : 'text-oxide-900' ]"></rect>
          </g>
        </svg>
      </div>
      <div class="flex flex-col gap-2">
        <div v-for="activity in weeklyActivity"
             :key="activity.timestamp"
             class="border border-oxide-900 p-4 flex space-between">
          <div class="text-left flex-grow font-mono">
            <div class="inline-block px-2 py-1 text-sm rounded-xs"
                 :class="[ colorize(activity) ]">{{ activity.type }}</div>
            <div class="mt-2">
              <NuxtLink :to="link(activity)"
                        class="text-gray-400 text-xs border-b border-transparent hover:border-oxide-400 hover:text-oxide-400">{{ activity.subject }}</NuxtLink>
            </div>
          </div>
          <div class="text-right text-sm text-gray-300 uppercase font-mono">{{ $moment(activity.timestamp).fromNow() }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapState } from "vuex";

export default {
  computed: {
    ...mapState("api", ["activities"]),
    ...mapGetters("api", ["reversedActivities"]),
    weeklyActivity() {
      return this.reversedActivities.slice().filter(item => this.$moment().diff(this.$moment(item.timestamp), 'week') < 1)
    },
    activitySummary() {
      // Assume first activity has first timestamp, last activity last
      let lastTimestamp = this.$moment();
      let firstTimestamp = lastTimestamp.clone().subtract("1", "year").endOf('week');
      let numberOfWeeks = lastTimestamp.diff(firstTimestamp, "week");
      let numberOfDaysPerWeek = 7;

      let weeks = Array(numberOfWeeks);
      for (var i = 0; i < weeks.length; i++) {
        if (i < weeks.length - 1) {
          weeks[i] = Array(numberOfDaysPerWeek).fill(0);
        } else {
          weeks[i] = Array(lastTimestamp.day() + 1).fill(0);
        }
      }

      this.activities.forEach((item) => {
        let ts = this.$moment(item.timestamp)
        let idx = ts.diff(firstTimestamp, "week")

        if (idx >= 0) weeks[idx-1][ts.day()]++
      })

      return weeks;
    },
  },
  mounted() {
    this.$store.dispatch("api/activities");
  },
  methods: {
    colorize(activity) {
      return {
        MACHINE_CREATED: ["bg-yellow-500"],
        MACHINE_POWERON: ["bg-oxide-700"],
        MACHINE_POWEROFF: ["bg-gray-500"],
        MACHINE_REBOOT: ["bg-yellow-500"],
        MACHINE_DELETED: ["bg-rod-600"],
      }[activity.type];
    },
    link(activity) {
      if (activity.type.startsWith("MACHINE_")) {
        return `/machines/${activity.subject}`;
      }
    },
  },
};
</script>
