<template>
  <transition name="modal-fade">
    <div v-if="show" class="modal">
      <div class="modal-dialog">
        <div class="dialog-header">
          <h1>添加用户</h1>
        </div>

        <div class="dialog-body">
          <form @submit="submit" class="dialog-form">

            <div class="form-control">
              <label for="email">邮箱地址</label>
              <div class="form-input"><input id="email" name="email" placeholder="邮箱地址" type="text"></div>
            </div>
            <div class="form-control">
              <label for="password">密码</label>
              <div class="form-input"><input id="password" name="password" placeholder="密码" type="password"></div>
            </div>
            <div class="form-control">
              <label for="role">用户组</label>
              <div class="form-input">
                <select id="role" name="role">
                  <option value="1">管理员账户</option>
                  <option value="2">普通用户</option>
                </select>
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
    show: false
  }),
  methods: {
    open() {
      this.show = true;
    },
    close() {
      this.show = false
    },
    submit(e) {
      let data = new FormData(e.target)

      this.$api.post("/users/create", data).then(res => {
        if (res.data.code && res.data.code === 20001) {
          let { msg } = res.data
          alert(msg.message)
        }

        location.reload()
      })

      return e.preventDefault()
    }
  }
}
</script>
