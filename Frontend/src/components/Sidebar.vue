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
            :label="item.text"
            @click="!item.options && redirectToPage(item.path)"
            @mouseover="openMenu(index)"
            @mouseout="closeMenu(index)"
          >
            <q-item-section style="font-size: 17px;">
              {{ item.text }}
            </q-item-section>

            <q-menu
              transition-show="jump-down"
              :ref="`menu-${index}`"
              :style="item.style"
              @mouseover="menuVisible = true"
              @mouseout="closeMenu(index)"
              @before-show="beforeShow(index)"
            >
              <q-list>
                <Options :options="item.options" />
              </q-list>
            </q-menu>
          </q-item>
        </q-list>
      </q-toolbar-title>

      <q-separator dark vertical />
      <q-separator dark vertical />

      <div class="q-pa-lg">
        <q-btn flat no-caps style="font-size: 17px;" @click="showLogin()">
          <svg class="q-mr-xs" height="28" width="28" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"> <path d="M12.12 12.78C12.05 12.77 11.96 12.77 11.88 12.78C10.12 12.72 8.71997 11.28 8.71997 9.50998C8.71997 7.69998 10.18 6.22998 12 6.22998C13.81 6.22998 15.28 7.69998 15.28 9.50998C15.27 11.28 13.88 12.72 12.12 12.78Z" stroke="#ffffff" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path> <path d="M18.74 19.3801C16.96 21.0101 14.6 22.0001 12 22.0001C9.40001 22.0001 7.04001 21.0101 5.26001 19.3801C5.36001 18.4401 5.96001 17.5201 7.03001 16.8001C9.77001 14.9801 14.25 14.9801 16.97 16.8001C18.04 17.5201 18.64 18.4401 18.74 19.3801Z" stroke="#ffffff" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path> <path d="M12 22C17.5228 22 22 17.5228 22 12C22 6.47715 17.5228 2 12 2C6.47715 2 2 6.47715 2 12C2 17.5228 6.47715 22 12 22Z" stroke="#ffffff" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path> </g></svg>
          Entrar
        </q-btn>

        <q-btn flat no-caps style="font-size: 17px;" @click="redirectToPage('/suporte')">
          <svg class="q-mr-xs" height="26" width="26" fill="#ffffff" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"><path d="M12 2C6.486 2 2 6.486 2 12v4.143C2 17.167 2.897 18 4 18h1a1 1 0 0 0 1-1v-5.143a1 1 0 0 0-1-1h-.908C4.648 6.987 7.978 4 12 4s7.352 2.987 7.908 6.857H19a1 1 0 0 0-1 1V18c0 1.103-.897 2-2 2h-2v-1h-4v3h6c2.206 0 4-1.794 4-4 1.103 0 2-.833 2-1.857V12c0-5.514-4.486-10-10-10z"></path></g></svg>
          Suporte
        </q-btn>
      </div>
    </q-toolbar>
  </q-header>
</template>
<script>
import Login from './LoginComponent.vue'
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
          href: 'vender',
          imgType: 5
        },
        {
          text: 'Gerenciar meus anúncios',
          href: 'meus-anuncios',
          imgType: 6
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
          options: this.optionsComprar,
          style: 'width: 13em;'
        },
        {
          text: 'Vender',
          path: '/vender',
          options: this.optionsVender,
          style: 'width: 15em;'
        },
        {
          text: 'Serviços',
          path: '/serviços',
          options: this.optionsServicos,
          style: 'width: 10em;'
        },
        {
          text: 'Noticias',
          path: '/noticias'
        },
        {
          text: 'Ajuda',
          path: '/ajuda',
          options: this.optionsAjuda,
          style: 'width: 16em;'
        }
      ]
    }
  },
  methods: {
    redirectToPage (path) {
      this.$router.push({ path })
    },

    // <Gambiarra> Lógica de visibilidade hover após o cursor
    // percorrer de um item até outro.
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
    },
    // </Gambiarra>
    showLogin () {
      this.$q.dialog({
        component: Login,
        parent: this
      })
    }
  }
}
</script>
