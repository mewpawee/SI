<template lang="html">
  <p>hi</p>
</template>

<script>
export default {
  auth: false,
  layout: 'admin',
  async middleware({ $auth, redirect }) {
    if ($auth.loggedIn) {
      const roles = await $auth.user.realm_access.roles
      if (!roles.includes('admin')) {
        return redirect('/')
      }
    } else {
      return redirect('/')
    }
  }
}
</script>
