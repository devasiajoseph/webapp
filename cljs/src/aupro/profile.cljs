(ns aupro.profile
   (:require [centipair.ui :as ui]
             [reagent.core :as r]
             [centipair.ajax :as ajax]
             [centipair.spa :as spa]
             [centipair.components.input :as input]
             [centipair.components.notifier :as notifier]
             [centipair.location :as location]
             [aupro.form :as form]
             [centipair.file :as file]))


(def profile-id (r/atom {:id "profile_id" :value 0}))
(def full-name (r/atom {:id "full_name" :type "text" :class "cfi" :placeholder "Enter Full Name" :label "Full name"}))
(def about (r/atom {:id "about" :type "text-area" :class "textarea textarea-bordered textarea-lg w-full max-w-xs" :placeholder "About" :label "About profile"}))
(def instagram (r/atom {:id "instagram" :type "text" :class "cfi" :placeholder "Instagram link" :label "Instagram link"}))
(def linkedin (r/atom {:id "linkedin" :type "text" :class "cfi" :placeholder "LinkedIn link" :label "LinkedIn link"}))
(def facebook (r/atom {:id "facebook" :type "text" :class "cfi" :placeholder "Facebook link" :label "Facebook link"}))
(def twitter (r/atom {:id "twitter" :type "text" :class "cfi" :placeholder "Twitter link" :label "Twitter link"}))
(def youtube (r/atom {:id "youtube" :type "text" :class "cfi" :placeholder "Youtube link" :label "Youtube link"}))
(def tiktok (r/atom {:id "tiktok" :type "text" :class "cfi" :placeholder "Tiktok link" :label "Tiktok link"}))
(def profile-pic (r/atom {:id "profile_pic" :value "" :url "/api/profile/upload-dp" :progress 0}))

(defn save-profile
  []
  (ajax/form-post
   "/api/profile" [profile-id
                   full-name
                   about
                   instagram
                   linkedin
                   facebook
                   twitter
                   youtube
                   tiktok
                   location/country]
   (fn [response]
     (spa/redirect (str "/profile/edit/" (:profile_id response))))))

(def save-profile-button (r/atom {:label "Save Profile" :on-click save-profile}))
(defn profile-form 
  []
  [:div
   [:div {:class (if (> (:value @profile-id) 0)  "" "hidden")}
    [:img {:src (:value @profile-pic)
           :class "mx-auto shadow-xl border-solid border-2 border-gray-300 w-44 rounded-lg"}]
    (file/file-input profile-pic)]
   (form/generate-form "Profile"
                       "Update profile details"
                       [full-name about location/country instagram linkedin facebook twitter youtube tiktok]
                       save-profile-button
                       [])])



(defn new-profile-form
  []
  (input/remote-select-options location/country)
  (ui/render profile-form "app"))



(defn get-profile
  [id]
(ajax/get-json
 (str "/api/profile/" id) nil
 (fn [response]
   (input/update-value-map [profile-id
                            full-name
                            about
                            instagram
                            linkedin
                            facebook
                            twitter
                            youtube
                            tiktok
                            location/country
                            profile-pic] response))))

(defn edit-profile-form
  [id]
  (get-profile id)
  (swap! profile-pic assoc :url (str "/api/profile/upload-dp/" id))
  (input/remote-select-options location/country)
  (ui/render profile-form "app"))



(defn profile-table 
  []
  [:div {:class "overflow-x-auto w-full"}
   [:table {:class "table w-full"}
    [:thead [:tr [:th ""] [:th "Name"] [:th "Action"]]]
    [:tbody [:tr 
             [:td [:div {:class "flex items-center space-x-3"}
                       [:div {:class "avatar"}
                        [:div {:class "mask mask-squircle w-12 h-12"}
                         [:img {:src "http://localhost:8081/static/uploads/images/c39b674f-72e4-4c71-b974-3f0e5bcb9376.webp", :alt "Avatar Tailwind CSS Component"}]]]]]
             [:td "Elon musk"]
             [:td [:button {:class "btn btn-ghost btn-xs"} "Edit"]]]]]]
  )


(defn render-profile-list
  [page]
  (ui/render profile-table "app")
  )