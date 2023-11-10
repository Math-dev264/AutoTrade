<template>
  <q-item
    v-for="(option, index) in options"
    :key="index"
    class="row"
    clickable
    style="height: 40px;"
  >
    <q-item-section avatar class="col-4">
      <template v-if="option.imgType === 1">
        <q-img style="height: 30px; width: 30px;" src="../assets/svg/iconCar.svg" />
      </template>
      <template v-if="option.imgType === 2">
        <q-img style="height: 30px; width: 30px;" src="../assets/svg/iconNewCar.svg" />
      </template>
      <template v-if="option.imgType === 3">
        <q-img style="height: 35px; width: 35px;" src="../assets/svg/iconBike.svg" />
      </template>
      <template v-if="option.imgType === 4">
        <q-img style="height: 35px; width: 35px;" src="../assets/svg/iconNewBike.svg" />
      </template>
      <template v-if="option.imgType === 5">
        <q-img style="height: 30px; width: 30px;" src="../assets/svg/iconSell.svg" />
      </template>
      <template v-if="option.imgType === 6">
        <q-img style="height: 32px; width: 32px;" src="../assets/svg/iconAnounces.svg" />
      </template>
    </q-item-section>

    <q-item-section
      @click="redirectToPage(option.href)"
      style="font-size: 15px;"
      class="col-8"
    >
      {{ option.text }}
    </q-item-section>
  </q-item>
</template>
<script>
export default {
  name: 'OptionsSidebar',

  props: {
    options: {
      type: Array,
      required: true
    }
  },

  methods: {
    redirectToPage (href) {
      window.location.hash = href
    }
  }
}
</script>
