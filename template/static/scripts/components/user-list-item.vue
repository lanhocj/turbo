<template>
  <div class="user-item">
    <ul class="user-item-info">
      <li class="index">{{ index + 1 }}</li>
      <li class="email">{{ email }}</li>
      <li class="role">{{ role }}</li>
      <li class="node-num">{{ nodeNum }}</li>
      <li class="locked">
        <a :class="[ locked ? 'locked' : '' ]" @click="changeLockState" href="javascript:void(0);">{{ locked ? '锁定' : '正常' }}</a>
      </li>
    </ul>

    <div class="group">
      <a @click="changePassword" href="javascript:void(0);">修改密码</a>
      <a @click="userSettings" href="javascript:void(0);">设置用户</a>
      <a @click="removeUser" href="javascript:void(0);">删除账户</a>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    index: Number,
    email: String,
    role: String,
    nodeNum: Number,
    locked: Boolean
  },
  methods: {
    changePassword() {
      console.log(this.email)
      this.$modalUserChangePassword({ email: this.email }).open()
    },
    changeLockState() {
      let data = new FormData()
      data.append("email", this.email)
      this.$api.post("/users/flushSetUserLockState", data).then(({data}) => {
        alert(data.message)
        location.reload()
      })
    },
    removeUser() {
      let email = this.email
      let data = new FormData()
      data.append("email", this.email)

      if (confirm(`确定删除 [${email}] 这个账户吗？`)) {
        this.$api.post("/users/remove", data).then(({ data }) => {
          alert(data.message)
          location.reload()
        })
      }
    },
    userSettings() {
      console.log(this.email)

      let modal = this.$modalUserSettings({ email: this.email })

      modal.getNodesWithUser(this.email).then(() => {
        modal.open()
      })
    }
  }
}
</script>
