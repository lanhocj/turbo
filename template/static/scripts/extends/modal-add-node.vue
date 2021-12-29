<template>
  <transition name="modal-fade">
    <div v-if="show" class="modal">
      <div class="modal-dialog">
        <div class="dialog-header">
          <h1>编辑节点服务器信息</h1>
        </div>

        <div class="dialog-body">
          <form @submit="submit" class="dialog-form">
            <div class="form-control">
              <label for="nodeName">服务器名称</label>
              <div class="form-input">
                <input id="nodeName" name="nodeName" placeholder="服务器名称" type="text">
              </div>
            </div>

            <div class="form-control-group">
              <div class="form-control">
                <label for="nodeAddr">IP / Domain 地址</label>
                <div class="form-input">
                  <input id="nodeAddr" name="nodeAddr" placeholder="域或IP地址" type="text">
                </div>
              </div>

              <div class="form-control" style="width: 5rem">
                <label for="nodePort">端口</label>
                <div class="form-input">
                  <input id="nodePort" name="nodePort" placeholder="端口号" type="number">
                </div>
              </div>
            </div>

            <div class="form-control">
              <label for="nodeTag">服务端 Tag</label>
              <div class="form-input">
                <input id="nodeTag" name="nodeTag" placeholder="管理标识" type="text">
                <span class="text-label">设置一个标识，只可以设置小写英文或数字，且必须以英文开始</span>
              </div>
            </div>

            <div class="form-control-group">
              <div class="form-control">
                <label for="clientAddr">客户端链接</label>
                <div class="form-input">
                  <input id="clientAddr" name="clientAddr" placeholder="域或IP地址" type="text">
                </div>
              </div>

              <div class="form-control" style="width: 5rem">
                <label for="clientPort">端口</label>
                <div class="form-input">
                  <input id="clientPort" name="clientPort" placeholder="443" type="number">
                </div>
              </div>
            </div>
            <div class="form-footer">
              <div class="form-control">
                <button class="form-button" type="submit">添加</button>
              </div>
              <div class="form-control">
                <a @click="close" class="form-button cancel" href="javascript:void(0);">取消</a>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  data: () => ({
    show: false,
    title: "",
    content: ""
  }),

  methods: {
    open() {
      this.show = true
    },
    close() {
      this.show = false
    },
    submit(event) {
      let data = new FormData(event.target);

      this.$api.post("/node", data).then(res => {
        if (res.status == 200) {
          location.reload()
        }
      })

      return event.preventDefault()
    }
  },
  mounted() {
    console.log("init..")
  }
}
</script>
