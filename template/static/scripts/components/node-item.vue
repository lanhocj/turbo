<template>
  <div class="node-item">
    <div :class="['icon', stateColorClass ]">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-server" viewBox="0 0 16 16">
        <path d="M1.333 2.667C1.333 1.194 4.318 0 8 0s6.667 1.194 6.667 2.667V4c0 1.473-2.985 2.667-6.667 2.667S1.333 5.473 1.333 4V2.667z"/>
        <path d="M1.333 6.334v3C1.333 10.805 4.318 12 8 12s6.667-1.194 6.667-2.667V6.334a6.51 6.51 0 0 1-1.458.79C11.81 7.684 9.967 8 8 8c-1.966 0-3.809-.317-5.208-.876a6.508 6.508 0 0 1-1.458-.79z"/>
        <path d="M14.667 11.668a6.51 6.51 0 0 1-1.458.789c-1.4.56-3.242.876-5.21.876-1.966 0-3.809-.316-5.208-.876a6.51 6.51 0 0 1-1.458-.79v1.666C1.333 14.806 4.318 16 8 16s6.667-1.194 6.667-2.667v-1.665z"/>
      </svg>
    </div>
    <div class="node-info">
      <h1 class="node-name" v-text="name">NodeName</h1>

      <ul class="node-extra">
        <li><p><span v-text="addr"></span></p></li>
        <li class="node-state"><span :class="stateColorClass" v-text="stateText"></span></li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    name: String,
    addr: {
      type: String,
    },
    port: String,
    state: String,
    tag: String
  },

  data: ()=>({
    available: -1
  }),
  created() {
    if (parseInt(this.state) === -1) {
      this.check()
    }
  },
  methods: {
    check() {
      let data = new FormData();
      data.append("addr", this.addr)
      data.append("port", this.port)

      this.$api.post("/node/available", data).then((res) => {
        if (res.status == 200) {
          this.available = res.data.state;
        }
      }).catch((err) => {
        console.log(err)
      })
    }
  },
  computed: {
    stateText() {
      return this.available == 0 ? "连接失败" : this.available == 1 ? "连接成功" : "测试中.."
    },
    stateColorClass() {
      return this.available == 0 ? "offline" : this.available == 1 ? "online" : "idle"
    }
  }

}
</script>
