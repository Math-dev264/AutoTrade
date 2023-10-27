<template>
  <q-header elevated>
    <q-toolbar style="background-color: #19191f; height: 6em;" class="no-padding">
      <div class="col-4 q-mr-lg">
        <q-img
          class="cursor-pointer"
          src="../assets/logo.png"
          style="width: 25em; height: 4em;"
          @click="redirectToPage('/home')"
        />
      </div>

      <q-toolbar-title class="q-pl-xl">
        <q-list padding class="flex q-gutter-sm">
          <q-item
            v-for="(item, index) in items"
            :key="index"
            v-ripple
            clickable
            active-class="text-info"
            :label="item.text"
            @click="!item.options && redirectToPage(item.path)"
            @mouseover="openMenu(index)"
            @mouseout="closeMenu(index)"
          >
            <q-item-section style="font-size: 17px;">
              {{ item.text }}
            </q-item-section>

            <q-menu
              :ref="`menu-${index}`"
              style="width: 13em;"
              @mouseover="menuVisible = true"
              @mouseout="closeMenu(index)"
              @before-show="beforeShow(index)"
            >
              <q-list separator>
                <Options :options="item.options" />
              </q-list>
            </q-menu>
          </q-item>
        </q-list>
      </q-toolbar-title>

      <q-separator dark vertical />
      <q-separator dark vertical />

      <div>
        <q-tabs v-model="tab2" shrink>
          <q-tab name="tab1" label="Entrar" icon="person" />
          <q-tab name="tab3" icon="chat" />
        </q-tabs>
      </div>
    </q-toolbar>
  </q-header>
</template>
<script>
import Options from './OptionsSidebar.vue'
export default {
  name: 'SideBar',
  components: { Options },
  data () {
    return {
      menuVisible: false,
      menuIndex: null,
      optionsComprar: [
        {
          text: 'Carros usados',
          href: 'carros-usados',
          imgType: 1
        },
        {
          text: 'Carros novos',
          href: 'carros-novos',
          imgType: 2
        },
        {
          text: 'Motos usadas',
          href: 'motos-usadas',
          imgType: 3
        },
        {
          text: 'Motos novos',
          href: 'motos-novos',
          imgType: 4
        }
      ],
      optionsVender: [
        {
          text: 'Vender veículo',
          href: 'vender'
        },
        {
          text: 'Gerenciar meus anúncios',
          href: 'meus-anuncios'
        }
      ],
      optionsServicos: [
        {
          text: 'Tabela FIPE',
          href: 'tabela-fipe'
        },
        {
          text: 'Financiamento',
          href: 'financiamento'
        },
        {
          text: 'Seguro veícular',
          href: 'seguro-veicular'
        }
      ],
      optionsAjuda: [
        {
          text: 'Central de ajuda',
          href: 'central-de-ajuda'
        },
        {
          text: 'Segurança',
          href: 'seguranca'
        },
        {
          text: 'Termos de uso e Política de privacidade',
          href: 'termos-de-uso-e-politica-de-privacidade'
        }
      ]
    }
  },
  computed: {
    items () {
      return [
        {
          text: 'Comprar',
          path: '/comprar',
          options: this.optionsComprar
        },
        {
          text: 'Vender',
          path: '/vender',
          options: this.optionsVender
        },
        {
          text: 'Serviços',
          path: '/serviços',
          options: this.optionsServicos
        },
        {
          text: 'Noticias',
          path: '/noticias'
        },
        {
          text: 'Ajuda',
          path: '/ajuda',
          options: this.optionsAjuda
        }
      ]
    }
  },
  methods: {
    redirectToPage (path) {
      this.$router.push({ path })
    },

    // <Gambiarra> Lógica para que fique invisível após o cursor
    // percorrer de um item até outro. (Hover)
    // Sujeita a mudanças e reformulação futura!
    openMenu (index) {
      this.menuVisible = true

      this.$refs[`menu-${index}`][0].show()
    },
    closeMenu (index) {
      this.menuVisible = false
      this.menuIndex = index

      setTimeout(() => {
        if (!this.menuVisible) {
          this.$refs[`menu-${index}`][0].hide()
        }
      }, 50)
    },
    beforeShow (index) {
      if (index !== this.menuIndex) {
        if (this.menuVisible && this.menuIndex !== null) {
          this.$refs[`menu-${this.menuIndex}`][0].hide()
        }
      }
    }
    // </Gambiarra>
  }
}
</script>
