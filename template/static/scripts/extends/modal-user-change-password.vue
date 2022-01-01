<template>
  <transition name="modal-fade">
    <div v-if="show" class="modal">
      <div class="modal-dialog">
        <div class="dialog-header">
          <h1>修改密码</h1>
        </div>

        <div class="dialog-body">
          <form @submit="submit" class="dialog-form">

            <div class="form-control">
              <label for="email">邮箱地址</label>
              <div class="form-input"><input readonly id="email" :value="email" name="email" placeholder="邮箱地址" type="text"></div>
            </div>
            <div class="form-control">
              <label for="password">密码</label>
              <div class="form-input"><input id="password" name="password" placeholder="密码" type="password"></div>
            </div>

            <div class="form-footer">
              <div class="form-control">
                <button class="form-button" type="submit">修改</button>
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
    email: ""
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

      this.$api.post("/users/change-password", data).then(({ data }) => {
        alert("密码修改成功")
        location.reload()
      })
      return e.preventDefault()
    }
  }
}
</script>
