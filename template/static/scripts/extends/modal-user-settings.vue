<template>
  <transition name="modal-fade">
    <div v-if="show" class="modal">
      <div class="modal-dialog">
        <div class="dialog-header">
          <h1>「{{ email }}」用户修改</h1>
        </div>

        <div class="dialog-body">
          <form @submit="submit" class="dialog-form">
            <input name="email" :value="email" hidden>
            <div class="form-control">
              <label for="role">用户组</label>
              <div class="form-input">
                <select v-model="roleId" id="role" name="role">
                  <option value="1">管理员账户</option>
                  <option value="2">普通用户</option>
                  <option value="0">锁定该用户</option>
                </select>
              </div>
            </div>

            <div class="form-control">
              <label>节点选择</label>

              <div class="node-select-box">
                <label  v-for="(node, index) in nodes" :key="index">
                  <input :checked="node.using" :value="node.id" type="checkbox" name="node">
                  <span class="checkbox-control">
                    {{ node.name }}
                  </span>
                </label>
              </div>
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
    email: '',
    roleId: 0,
    nodes: []
  }),
  created() {
    // this.getNodesWithUser(this.email)
  },
  methods: {
    getNodesWithUser(email, roleId = 0) {
      this.roleId = roleId
      let data = new FormData()
      data.append("email", email)
      return this.$api.post("/users/nodes", data).then((res) => {
        return res.data
      }).then(data => {
        this.nodes = data
      })
    },
    open() {
      this.show = true;
    },
    close() {
      this.show = false
    },
    submit(e) {
      let data = new FormData(e.target)
      let nodes = data.getAll("node")

      data.append("node", nodes)

      this.$api.post("/users/setting", data).then(res => {
        return res.data
      }).then(r => {
        if (r.status == 2000) {
          alert(r.message)
        }
        location.reload()
      })

      return e.preventDefault()
    }
  }
}
</script>
