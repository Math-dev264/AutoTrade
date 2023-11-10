<template>
  <q-layout view="lHh Lpr lFf" container style="height: 100vh; width: 100%;">
    <Sidebar />

    <q-page-container>
      <q-page>
        <router-view />
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script>
import Sidebar from 'src/components/Sidebar.vue'
export default {
  name: 'MainLayout',
  components: { Sidebar }
}
</script>
