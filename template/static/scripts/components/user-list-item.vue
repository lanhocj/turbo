<template>
  <div class="user-item">
    <ul class="user-item-info">
      <li class="index">{{ index + 1 }}</li>
      <li :class="['email', { locked : !roleId }]">
        <ul>
          <li class="profile">
            {{ email }}
            <span v-if="roleId == 1" class="roleName">
            <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" fill="currentColor" class="bi bi-patch-check-fill" viewBox="0 0 16 16">
              <path d="M10.067.87a2.89 2.89 0 0 0-4.134 0l-.622.638-.89-.011a2.89 2.89 0 0 0-2.924 2.924l.01.89-.636.622a2.89 2.89 0 0 0 0 4.134l.637.622-.011.89a2.89 2.89 0 0 0 2.924 2.924l.89-.01.622.636a2.89 2.89 0 0 0 4.134 0l.622-.637.89.011a2.89 2.89 0 0 0 2.924-2.924l-.01-.89.636-.622a2.89 2.89 0 0 0 0-4.134l-.637-.622.011-.89a2.89 2.89 0 0 0-2.924-2.924l-.89.01-.622-.636zm.287 5.984-3 3a.5.5 0 0 1-.708 0l-1.5-1.5a.5.5 0 1 1 .708-.708L7 8.793l2.646-2.647a.5.5 0 0 1 .708.708z"/>
            </svg>
          </span>
          </li>
          <li v-if="roleId" class="hash">
            <span>{{ token ? token : '-' }}</span>
            <a @click="flushToken" class="icon" href="javascript:void(0);">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-arrow-repeat" viewBox="0 0 16 16">
                <path d="M11.534 7h3.932a.25.25 0 0 1 .192.41l-1.966 2.36a.25.25 0 0 1-.384 0l-1.966-2.36a.25.25 0 0 1 .192-.41zm-11 2h3.932a.25.25 0 0 0 .192-.41L2.692 6.23a.25.25 0 0 0-.384 0L.342 8.59A.25.25 0 0 0 .534 9z"/>
                <path fill-rule="evenodd" d="M8 3c-1.552 0-2.94.707-3.857 1.818a.5.5 0 1 1-.771-.636A6.002 6.002 0 0 1 13.917 7H12.9A5.002 5.002 0 0 0 8 3zM3.1 9a5.002 5.002 0 0 0 8.757 2.182.5.5 0 1 1 .771.636A6.002 6.002 0 0 1 2.083 9H3.1z"/>
              </svg>

              刷新
            </a>
          </li>
        </ul>
      </li>
      <li class="node-num">{{ nodeNum }}</li>
    </ul>

    <div class="group">
      <a @click="toggleDropdownMenus" href="javascript:void(0);">设置</a>

      <ul v-if="dropdownMenu" ref="dropdownMenu" class="dropdown-menus">
        <li><a @click="changePassword" href="javascript:void(0);">修改密码</a></li>
        <li><a @click="userSettings" href="javascript:void(0);">节点管理</a></li>
        <li><a @click="removeUser" href="javascript:void(0);">删除</a></li>
      </ul>
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
    locked: Boolean,
    token: String,
    roleId: Number
  },
  data: () =>({
    dropdownMenu: false
  }),
  methods: {
    toggleDropdownMenus() {
      this.dropdownMenu = !this.dropdownMenu
    },
    flushToken() {
      let email = this.email
      let data = new FormData();
      data.append("email", email)

      if (confirm("确定需要更新订阅地址吗？")) {
        this.$api.post("users/token-refresh", data).then(({data}) => {
          location.reload()
        })
      }
    },
    changePassword() {
      if (this.dropdownMenu) {
        this.dropdownMenu = false
      }

      this.$modalUserChangePassword({ email: this.email }).open()
    },
    changeLockState() {
      let data = new FormData()
      data.append("email", this.email)
      this.$api.post("/users/lock", data).then(({data}) => {
        alert(data.message)
        location.reload()
      })
    },
    removeUser() {
      if (this.dropdownMenu) {
        this.dropdownMenu = false
      }

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
      if (this.dropdownMenu) {
        this.dropdownMenu = false
      }

      console.log(this.email)

      let modal = this.$modalUserSettings({ email: this.email })

      modal.getNodesWithUser(this.email, this.roleId).then(() => {
        modal.open()
      })
    }
  },
  mounted() {
    document.addEventListener('click', (e) => {
      if (!this.$el.contains(e.target)) this.dropdownMenu = false
    })
  }
}
</script>
