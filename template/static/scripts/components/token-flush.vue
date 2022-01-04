<template>
  <section class="section">
    <div class="section-header">
      <h1 class="title">订阅信息</h1>
    </div>

    <div class="section-body">
      <div class="view">
        <h1>手机扫描二维码</h1>
        <div>
          <img :src="qrcodeUrl" alt="">
        </div>
      </div>

      <hr />


      <div class="form-control">
        <div class="form-input form-group">
          <input readonly :value="currentUrl" type="text">

          <button @click="copy" class="copy">一键复制</button>
        </div>
      </div>

      <p>如果订阅不成功，请<a @click="flushToken" href="javascript:void(0)">刷新订阅</a>地址再次尝试</p>
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
        this.$api.post("users/token-refresh", data).then(({data}) => {
          this.currentUrl = data.url
          location.reload()
        })
      }
    },
    copy(e) {
      let transfer = document.createElement('input');
      document.body.appendChild(transfer);
      transfer.value = this.currentUrl;
      transfer.focus();
      transfer.select();
      if (document.execCommand('copy')) {
        document.execCommand('copy');
      }
      transfer.blur();
      console.log('复制成功');
      document.body.removeChild(transfer);
    }
  },
  computed: {
    qrcodeUrl() {
      let url = window.btoa(this.currentUrl)
      let name = window.encodeURI("Turbo 代理")
      let data = `sub://${url}#${name}`
      return `https://api.qrserver.com/v1/create-qr-code/?size=150x150&data=${data}`
    }
  }
}
</script>