<template>
  <div class="user-item">
    <ul class="user-item-info">
      <li class="index">{{ index + 1 }}</li>
      <li class="email">{{ email }}</li>
      <li class="role">{{ role }}</li>
      <li class="node-num">{{ nodeNum }}</li>
      <li class="locked">{{ locked ? '锁定' : '正常' }}</li>
    </ul>

    <div class="group">
      <a @click="changePassword" href="javascript:void(0);">修改密码</a>
      <a @click="userSettings" href="javascript:void(0);">设置用户</a>
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
