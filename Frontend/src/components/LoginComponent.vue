<template>
  <q-dialog ref="dialog">
    <div class="col-12 col-md-12" style="width: 60%;">
      <q-card>
        <div class="col-12 row" style="height: 80vh;">
          <div class="col-6 my-component">
            <div class="background-image" />

            <q-card-section style="height: 100%; padding-top: 7em; font-family: cursive;">
              <div class="col-12 row justify-center">
                <span class="text-h4 text-white">Seja bem vindo ao:</span>
              </div>

              <div class="col-12 row justify-center">
                <q-img
                  src="../assets/logo.png"
                  style="width: 27em; height: 9em;"
                />
              </div>

              <div class="col-12 row q-px-lg text-h6 text-white" >
                <span>Aqui você pode encontrar o veículo que desejar para que enfim, você possa entrar em contato com um vendedor com fins de negociar a compra de um veículo desejado. Além disso, você também pode negociar possíveis trocas e vendas de seus veículos.</span>
                <br />
                <span>Para utilizar a plataforma, você pode <a class="cursor-pointer" style="color: #04ffff;" @click="type = 'login'">entrar</a> com sua conta já existente ou então <a class="cursor-pointer" style="color: #04ffff;" @click="type = 'signup'">criar uma nova.</a></span>
              </div>

              <br /><br />
              <div class="col-12 row justify-center" style="font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;">
                <q-btn outline class="text-white" label="Clique aqui para saber mais" />
              </div>
            </q-card-section>
          </div>
          <div class="col-6">
            <q-card-section class="q-pl-xl">
              <div class="col-12 row justify-end">
                <q-btn v-close-popup flat icon="close" />
              </div>

              <br /><br />
              <div class="col-12 row text-h4">
                <span v-if="type === 'login'">
                  Entrar
                </span>

                <span v-if="type === 'signup'">
                  Criar uma conta
                </span>
              </div>
            </q-card-section>

            <template v-if="type === 'login'">
              <q-card-section class="q-px-xl">
                <q-form class="q-gutter-md row">
                  <q-input
                    outlined
                    v-model="email"
                    label="E-Mail"
                    class="col-12"
                  />
                  <q-input
                    outlined
                    v-model="senha"
                    label="Senha"
                    class="col-12"
                    :type="isPwdVisible ? 'text' : 'password'"
                  >
                    <template #append>
                      <q-icon
                        :name="isPwdVisible ? 'visibility' : 'visibility_off'"
                        class="cursor-pointer"
                        @click="isPwdVisible = !isPwdVisible"
                      />
                    </template>
                  </q-input>
                </q-form>
              </q-card-section>

              <q-card-section class="row justify-between q-pt-sm q-pl-xl">
                <q-btn label="entrar" style="width: 8em; background-color: #2c9894; color: #fff;" />
                <span class="q-pr-md text-bold cursor-pointer" style="color: #2c9894;">Esqueceu a senha?</span>
              </q-card-section>

              <q-card-section class="row q-mx-lg items-center">
                <div class="col-5">
                  <q-separator />
                </div>

                <div class="col-2 row justify-center">
                  <span class="text-h5">Ou</span>
                </div>

                <div class="col-5">
                  <q-separator />
                </div>
              </q-card-section>

              <q-card-section class="q-px-xl row q-gutter-md">
                <q-btn outline class="col-12" style="color: #2c9894;">
                  <svg height="26" width="26" enable-background="new 0 0 512 512" id="Layer_1" version="1.1" viewBox="0 0 512 512" xml:space="preserve" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><g><path d="M42.4,145.9c15.5-32.3,37.4-59.6,65-82.3c37.4-30.9,80.3-49.5,128.4-55.2c56.5-6.7,109.6,4,158.7,33.4   c12.2,7.3,23.6,15.6,34.5,24.6c2.7,2.2,2.4,3.5,0.1,5.7c-22.3,22.2-44.6,44.4-66.7,66.8c-2.6,2.6-4,2.4-6.8,0.3   c-64.8-49.9-159.3-36.4-207.6,29.6c-8.5,11.6-15.4,24.1-20.2,37.7c-0.4,1.2-1.2,2.3-1.8,3.5c-12.9-9.8-25.9-19.6-38.7-29.5   C72.3,169,57.3,157.5,42.4,145.9z" fill="#E94335"/><path d="M126,303.8c4.3,9.5,7.9,19.4,13.3,28.3c22.7,37.2,55.1,61.1,97.8,69.6c38.5,7.7,75.5,2.5,110-16.8   c1.2-0.6,2.4-1.2,3.5-1.8c0.6,0.6,1.1,1.3,1.7,1.8c25.8,20,51.7,40,77.5,60c-12.4,12.3-26.5,22.2-41.5,30.8   c-43.5,24.8-90.6,34.8-140.2,31C186.3,501.9,133,477.5,89,433.5c-19.3-19.3-35.2-41.1-46.7-66c10.7-8.2,21.4-16.3,32.1-24.5   C91.6,329.9,108.8,316.9,126,303.8z" fill="#34A853"/><path d="M429.9,444.9c-25.8-20-51.7-40-77.5-60c-0.6-0.5-1.2-1.2-1.7-1.8c8.9-6.9,18-13.6,25.3-22.4   c12.2-14.6,20.3-31.1,24.5-49.6c0.5-2.3,0.1-3.1-2.2-3c-1.2,0.1-2.3,0-3.5,0c-40.8,0-81.7-0.1-122.5,0.1c-4.5,0-5.5-1.2-5.4-5.5   c0.2-29,0.2-58,0-87c0-3.7,1-4.7,4.7-4.7c74.8,0.1,149.6,0.1,224.5,0c3.2,0,4.5,0.8,5.3,4.2c6.1,27.5,5.7,55.1,2,82.9   c-3,22.2-8.4,43.7-16.7,64.5c-12.3,30.7-30.4,57.5-54.2,80.5C431.6,443.8,430.7,444.3,429.9,444.9z" fill="#4285F3"/><path d="M126,303.8c-17.2,13.1-34.4,26.1-51.6,39.2c-10.7,8.1-21.4,16.3-32.1,24.5C34,352.1,28.6,335.8,24.2,319   c-8.4-32.5-9.7-65.5-5.1-98.6c3.6-26,11.1-51,23.2-74.4c15,11.5,29.9,23.1,44.9,34.6c12.9,9.9,25.8,19.7,38.7,29.5   c-2.2,10.7-5.3,21.2-6.3,32.2c-1.8,20,0.1,39.5,5.8,58.7C125.8,301.8,125.9,302.8,126,303.8z" fill="#FABB06"/></g></svg>
                  <span class="col-6 text-black">Entrar com o Google</span>
                </q-btn>
                <q-btn outline class="col-12" style="color: #2c9894;">
                  <svg height="26" width="26" xmlns="http://www.w3.org/2000/svg" aria-label="Facebook" role="img" viewBox="0 0 512 512"><rect width="512" height="512" rx="15%" fill="#1877f2"/><path d="M355.6 330l11.4-74h-71v-48c0-20.2 9.9-40 41.7-40H370v-63s-29.3-5-57.3-5c-58.5 0-96.7 35.4-96.7 99.6V256h-65v74h65v182h80V330h59.6z" fill="#fff"/></svg>
                  <span class="col-6 text-black q-pl-md">Entrar com o Facebook</span>
                </q-btn>
                <q-btn outline class="col-12" style="color: #2c9894;">
                  <svg height="26" width="26" style="shape-rendering:geometricPrecision; text-rendering:geometricPrecision; image-rendering:optimizeQuality; fill-rule:evenodd; clip-rule:evenodd" version="1.1" viewBox="0 0 201 201" xml:space="preserve" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><defs></defs><g id="Layer_x0020_1"><g id="_541895288"><rect fill="#1877f2" height="201" rx="19" ry="19" width="201"/><g><polygon fill="#FEFEFE" points="39,102 39,162 68,162 68,97 68,80 39,80 "/><path fill="#FEFEFE" d="M54 39c-8,0 -15,7 -15,15 0,7 7,14 15,14 7,0 14,-7 14,-14 0,-8 -7,-15 -14,-15z"/><path fill="#FEFEFE" d="M161 105c-2,-15 -9,-25 -30,-25 -12,0 -20,5 -24,11l0 0 0 -11 -23 0 0 16 0 66 24 0 0 -41c0,-10 2,-21 15,-21 13,0 14,13 14,22l0 40 25 0 0 -45 0 0c0,-4 0,-8 -1,-12z"/></g></g></g></svg>
                  <span class="col-6 text-black q-pl-sm">Entrar com o Linkedin</span>
                </q-btn>
              </q-card-section>

              <br />
              <q-card-section class="row justify-center q-px-xl">
                <span class="q-pr-xs">Não possui uma conta?</span>
                <span class="text-bold cursor-pointer" style="color: #2c9894;" @click="type = 'signup'">Criar conta</span>
              </q-card-section>
            </template>

            <template v-if="type === 'signup'">
              <q-card-section class="q-px-xl">
                <q-form class="q-gutter-md row">
                  <q-input
                    outlined
                    v-model="nome"
                    label="Nome Completo"
                    class="col-12"
                  />
                  <q-input
                    outlined
                    v-model="email"
                    label="E-Mail"
                    class="col-12"
                  />
                  <q-input
                    outlined
                    v-model="senha"
                    label="Senha"
                    class="col-12"
                    :type="isPwdVisible ? 'text' : 'password'"
                  >
                    <template #append>
                      <q-icon
                        :name="isPwdVisible ? 'visibility' : 'visibility_off'"
                        class="cursor-pointer"
                        @click="isPwdVisible = !isPwdVisible"
                      />
                    </template>
                  </q-input>
                </q-form>
              </q-card-section>

              <q-card-section class="row justify-between q-pt-sm q-pl-xl">
                <q-btn label="criar conta" style="width: 12em; background-color: #2c9894; color: #fff;" />
              </q-card-section>

              <q-card-section class="row q-mx-lg items-center">
                <div class="col-5">
                  <q-separator />
                </div>

                <div class="col-2 row justify-center">
                  <span class="text-h5">Ou</span>
                </div>

                <div class="col-5">
                  <q-separator />
                </div>
              </q-card-section>

              <q-card-section class="q-px-xl row q-gutter-md">
                <q-btn outline class="col-12" style="color: #2c9894;">
                  <svg height="26" width="26" enable-background="new 0 0 512 512" id="Layer_1" version="1.1" viewBox="0 0 512 512" xml:space="preserve" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><g><path d="M42.4,145.9c15.5-32.3,37.4-59.6,65-82.3c37.4-30.9,80.3-49.5,128.4-55.2c56.5-6.7,109.6,4,158.7,33.4   c12.2,7.3,23.6,15.6,34.5,24.6c2.7,2.2,2.4,3.5,0.1,5.7c-22.3,22.2-44.6,44.4-66.7,66.8c-2.6,2.6-4,2.4-6.8,0.3   c-64.8-49.9-159.3-36.4-207.6,29.6c-8.5,11.6-15.4,24.1-20.2,37.7c-0.4,1.2-1.2,2.3-1.8,3.5c-12.9-9.8-25.9-19.6-38.7-29.5   C72.3,169,57.3,157.5,42.4,145.9z" fill="#E94335"/><path d="M126,303.8c4.3,9.5,7.9,19.4,13.3,28.3c22.7,37.2,55.1,61.1,97.8,69.6c38.5,7.7,75.5,2.5,110-16.8   c1.2-0.6,2.4-1.2,3.5-1.8c0.6,0.6,1.1,1.3,1.7,1.8c25.8,20,51.7,40,77.5,60c-12.4,12.3-26.5,22.2-41.5,30.8   c-43.5,24.8-90.6,34.8-140.2,31C186.3,501.9,133,477.5,89,433.5c-19.3-19.3-35.2-41.1-46.7-66c10.7-8.2,21.4-16.3,32.1-24.5   C91.6,329.9,108.8,316.9,126,303.8z" fill="#34A853"/><path d="M429.9,444.9c-25.8-20-51.7-40-77.5-60c-0.6-0.5-1.2-1.2-1.7-1.8c8.9-6.9,18-13.6,25.3-22.4   c12.2-14.6,20.3-31.1,24.5-49.6c0.5-2.3,0.1-3.1-2.2-3c-1.2,0.1-2.3,0-3.5,0c-40.8,0-81.7-0.1-122.5,0.1c-4.5,0-5.5-1.2-5.4-5.5   c0.2-29,0.2-58,0-87c0-3.7,1-4.7,4.7-4.7c74.8,0.1,149.6,0.1,224.5,0c3.2,0,4.5,0.8,5.3,4.2c6.1,27.5,5.7,55.1,2,82.9   c-3,22.2-8.4,43.7-16.7,64.5c-12.3,30.7-30.4,57.5-54.2,80.5C431.6,443.8,430.7,444.3,429.9,444.9z" fill="#4285F3"/><path d="M126,303.8c-17.2,13.1-34.4,26.1-51.6,39.2c-10.7,8.1-21.4,16.3-32.1,24.5C34,352.1,28.6,335.8,24.2,319   c-8.4-32.5-9.7-65.5-5.1-98.6c3.6-26,11.1-51,23.2-74.4c15,11.5,29.9,23.1,44.9,34.6c12.9,9.9,25.8,19.7,38.7,29.5   c-2.2,10.7-5.3,21.2-6.3,32.2c-1.8,20,0.1,39.5,5.8,58.7C125.8,301.8,125.9,302.8,126,303.8z" fill="#FABB06"/></g></svg>
                  <span class="col-6 text-black">Criar com o Google</span>
                </q-btn>
                <q-btn outline class="col-12" style="color: #2c9894;">
                  <svg height="26" width="26" xmlns="http://www.w3.org/2000/svg" aria-label="Facebook" role="img" viewBox="0 0 512 512"><rect width="512" height="512" rx="15%" fill="#1877f2"/><path d="M355.6 330l11.4-74h-71v-48c0-20.2 9.9-40 41.7-40H370v-63s-29.3-5-57.3-5c-58.5 0-96.7 35.4-96.7 99.6V256h-65v74h65v182h80V330h59.6z" fill="#fff"/></svg>
                  <span class="col-6 text-black q-pl-md">Criar com o Facebook</span>
                </q-btn>
                <q-btn outline class="col-12" style="color: #2c9894;">
                  <svg height="26" width="26" style="shape-rendering:geometricPrecision; text-rendering:geometricPrecision; image-rendering:optimizeQuality; fill-rule:evenodd; clip-rule:evenodd" version="1.1" viewBox="0 0 201 201" xml:space="preserve" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><defs></defs><g id="Layer_x0020_1"><g id="_541895288"><rect fill="#1877f2" height="201" rx="19" ry="19" width="201"/><g><polygon fill="#FEFEFE" points="39,102 39,162 68,162 68,97 68,80 39,80 "/><path fill="#FEFEFE" d="M54 39c-8,0 -15,7 -15,15 0,7 7,14 15,14 7,0 14,-7 14,-14 0,-8 -7,-15 -14,-15z"/><path fill="#FEFEFE" d="M161 105c-2,-15 -9,-25 -30,-25 -12,0 -20,5 -24,11l0 0 0 -11 -23 0 0 16 0 66 24 0 0 -41c0,-10 2,-21 15,-21 13,0 14,13 14,22l0 40 25 0 0 -45 0 0c0,-4 0,-8 -1,-12z"/></g></g></g></svg>
                  <span class="col-6 text-black q-pl-sm">Criar com o Linkedin</span>
                </q-btn>
              </q-card-section>

              <q-card-section class="row justify-center q-px-xl">
                <span class="q-pr-xs">Já possui uma conta?</span>
                <span class="text-bold cursor-pointer" style="color: #2c9894;" @click="type = 'login'">Entrar</span>
              </q-card-section>
            </template>
          </div>
        </div>
      </q-card>
    </div>
  </q-dialog>
</template>
<script>
import { ref } from 'vue'

export default {
  setup () {
    const type = ref('login')
    const isPwdVisible = ref(false)

    return {
      type,
      isPwdVisible
    }
  },
  data () {
    return {
      email: '',
      senha: ''
    }
  }
}
</script>
<style>
.my-component {
  position: relative;
}
.background-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-image: url('../assets/fundoLogin.jpg');
  background-size: cover;
  background-position: center;
  z-index: 0;
}
.background-image::before {
  content: "";
  background-color: rgba(0, 141, 151, 0.829);
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}
.content {
  text-align: center;
  color: #fff;
  z-index: 1;
}
</style>
