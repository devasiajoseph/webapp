(ns centipair.db)

(def api {:user-login "/api/user/login"
          :user-logout "/api/user/logout"
          :user-status "/api/user/status"
          :user-profile "/api/user/profile"
          :change-my-password "/api/user/my-password"
          :user-list "/api/user/list"
          :user-delete "/api/user/delete"
          :user-create "/api/user/create"
          :user-update "/api/user/update"
          :user-edit "/api/user/edit"
          :change-user-password "/api/user/password"
          :save-base-file "/api/write/base"
          :save-static-file "/api/write/static"
          :save-template "/api/write/template"
          :save-markdown "/api/write/markdown"
          :save-dynamic "/api/write/dynamic"
          :pages "/api/pages"
          :search-pages "/api/search/pages"
          :search-images "/api/admin/search/images"
          :base-options "/api/pages/options/base"
          :template-options "/api/pages/options/template"
          :domain "/api/domain"
          :images "/api/admin/images"
          :delete-selected-images "/api/admin/delete/images"
          :open-file "/api/admin/files/open"
          :save-file "/api/admin/files/save"
          })

(defn api-paginate
  [k page]
  (str (k api) "?page=" page))
