<template>
  <section class="section">
    <div class="section-header">
      <h1 class="title">订阅信息</h1>
    </div>

    <div class="section-body">
      <div class="form-control">
        <div class="form-input">
          <input readonly :value="currentUrl" type="text">
        </div>
      </div>

      <p>如果订阅不成功，请<a @click="flushToken" href="javascript:void(0)">刷新订阅</a>地址再次上市</p>
    </div>
  </section>
</template>

<script>
export default {
  props: {
    url: String,
    email: String
  },
  data: () => ({
    currentUrl: ""
  }),
  created() {
    this.currentUrl = this.url
  },
  methods: {
    flushToken() {
      let email = this.email
      let data = new FormData();
      data.append("email", email)

      if (confirm("确定需要更新订阅地址吗？")) {
        this.$api.post("users/flushToken", data).then(({data}) => {
          this.currentUrl = data.url
          location.reload()
        })
      }
    }
  }
}
</script>