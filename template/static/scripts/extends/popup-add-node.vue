<template>
  <div v-show="show" class="popup">
    <h1>hello</h1>
  </div>
</template>

<script>
export default {
  data() {
    return {
      show: true,
      title: "",
      content: ""
    }
  },
  methods: {
    toggle() {
      this.show = !this.show
    }
  }
}
</script>
