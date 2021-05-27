export default async function({ $auth, redirect }) {
  if ($auth.loggedIn) {
    try {
      const roles = await $auth.user.realm_access.roles
      if (roles.includes('admin')) {
        return redirect('/admin')
      }
    } catch (err) {
      console.log(err)
    }
  }
}
