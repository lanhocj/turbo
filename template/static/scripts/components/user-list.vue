<template>
  <div class="user-list">
    <div class="user-item-header">
      <li class="index">ID</li>
      <li class="email">邮箱地址</li>
      <li class="role">权限</li>
      <li class="node-num">节点数量</li>
      <li class="locked">状态</li>
    </div>
    <user-list-item v-for="(item, index) in list" :index="index" :email="item.email" :node-num="item.nodeNum" :role="item.role" />
  </div>
</template>

<script>
export default {
  components: {
    UserListItem: () => import("./user-list-item")
  },
  data: () => ({
    list: []
  }),
  beforeCreate() {
    this.$api.get("/users").then(res => {
      if (res.status === 200) {
        console.log(res.data)
        this.list = res.data
      }
    })
  }
}
</script>
