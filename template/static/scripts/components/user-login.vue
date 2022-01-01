<template>
  <form @submit="submit" autocomplete="off" class="login-form" method="post">
    <div class="form-control">
      <label for="email">邮箱</label>
      <div class="form-input">
        <input id="email" placeholder="输入注册邮箱" name="email" type="text">
      </div>
    </div>

    <div class="form-control">
      <label for="password">密码</label>
      <div class="form-input">
        <input id="password" placeholder="输入管理密码" name="password" type="password">
      </div>
    </div>

    <div class="form-control form-footer">
      <button class="form-button" type="submit">登录账户</button>
    </div>
  </form>
</template>

<script>
export default {
  methods: {
    submit(e) {
      let data = new FormData(e.target)

      this.$api.post("/login", data).then(({ data }) => {
        if (data.message) {
          alert(data.message)
        }

        location.href = "/"
      })

      return e.preventDefault()
    }
  }
}
</script>
