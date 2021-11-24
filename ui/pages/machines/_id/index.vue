<template>
  <div class="flex flex-col gap-8">
    <subgroup name="CPUs">
      <canvas class="w-full h-32" id="cpu-timeline" />
    </subgroup>
    <subgroup name="Memory">
      <canvas class="w-full h-32" id="memory-timeline" />
    </subgroup>
    <subgroup name="Block devices">
      <canvas class="w-full h-32" id="block-timeline" />
    </subgroup>
  </div>
</template>

<script>
export default {
  data () {
    return {
      cpu: [],
      memory: [],
      block: [],
    }
  },
  mounted() {
    let last = 50;
    for (var i = 0; i < 100; i++) {
      let next = Math.min(100, Math.max(0, -10 + Math.random() * 20 + last))
      this.cpu.push(next)
      last = next
    }
    for (var i = 0; i < 100; i++) {
      let next = Math.min(100, Math.max(0, -10 + Math.random() * 20 + last))
      this.memory.push(next)
      last = next
    }
    for (var i = 0; i < 100; i++) {
      let next = Math.min(100, Math.max(0, -10 + Math.random() * 20 + last))
      this.block.push(next)
      last = next
    }
    this.draw('cpu-timeline', this.cpu)
    this.draw('memory-timeline', this.memory)
    this.draw('block-timeline', this.block)
  },
  methods: {
    draw (target, data) {
      let canvas = document.getElementById(target)
      let width = canvas.offsetWidth
      let height = canvas.offsetHeight
      canvas.width = width
      canvas.height = height

      let ctx = canvas.getContext('2d')

      ctx.strokeStyle = '#fff'
      ctx.clearRect(0, 0, width, height)

      // draw index on left side

      let segmentSpace = 4
      let segmentSize = 8 
      let numSegmentsX = width / (segmentSize + segmentSpace) - 1
      let numSegmentsY = height / (segmentSize + segmentSpace) - 1
      
      console.log(numSegmentsX, numSegmentsY)

      ctx.fillStyle = '#479073'
      for (var i = 0; i < numSegmentsX; i++) {
        // Compute target value range
        let start = Math.floor(i / numSegmentsX * data.length)
        let end = Math.min(data.length, Math.floor((i+1) / numSegmentsX * data.length))

        console.log(i, start, end, value)

        // Compute max value over range
        var value = 0
        for (var j = start; j <= Math.min(data.length - 1, end); j++) {
          value = Math.max(data[j], value)
        }


        // Compute how many vertical rects to fill
        let vertical = value / 100.0 * numSegmentsY 
        for (var j = 0; j < vertical; j++) {
          ctx.fillRect(i * (segmentSpace + segmentSize), height - j * (segmentSize + segmentSpace), segmentSize, segmentSize)
        }
      }
    }
  }
}
</script>