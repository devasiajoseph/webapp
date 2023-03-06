(ns aupro.feed
  (:require [reagent.core :as r]
            [centipair.ui :as ui]))


(def post (r/atom {}))



(defn post-ui 
  []
  [:div {:class "post-box"}
   
   [:div {:class "flex flex-shrink-0 p-4 pb-0"}
    [:a {:href (:profile-link post) :class "flex-shrink-0 group block"}
     [:div {:class "flex items-top"}
      [:div
       [:img {:class "inline-block h-9 w-9 rounded-full", :src "https://pbs.twimg.com/profile_images/1308769664240160770/AfgzWVE7_normal.jpg"}]]
      [:div {:class "ml-3"}
       [:p {:class "flex items-center text-base leading-6 font-medium text-gray-800"} "Joe Biden\n                "
        [:span {:class "ml-1 text-sm leading-5 font-medium text-gray-400 group-hover:text-gray-300 transition ease-in-out duration-150"} "@JoeBiden . Nov 7 2022"]]]]]]
   [:div {:class "pl-16"}
    [:p {:class "text-base width-auto font-medium text-gray-800 flex-shrink"} "America, I’m honored that you have chosen me to lead our great\n            country.\n            The work ahead of us will be hard, but I promise you this: I\n            will be a"
     [:a {:href "#", :class "text-blue-400 hover:underline"} "#President"] "for all Americans — whether you voted for me or not.\n            I will keep the faith that you have placed in me."]
    [:div {:class "flex my-3 mr-2 rounded-2xl border border-gray-600"}
     [:img {:class "rounded-2xl", :src "https://pbs.twimg.com/media/EnTkhz-XYAEH2kY?format=jpg&name=small"}]]
    [:div {:class "flex"}
     [:div {:class "w-full"}
      [:div {:class "flex items-center"}
       [:div {:class "flex-1 flex items-center text-gray-800 text-xs text-gray-400 hover:text-blue-400 dark:hover:text-blue-400 transition duration-350 ease-in-out"}
      "12.3 k"]
       [:div {:class "flex-1 flex items-center text-gray-800 text-xs text-gray-400 hover:text-green-400 dark:hover:text-green-400 transition duration-350 ease-in-out"}
        "14 k"]
       [:div {:class "flex-1 flex items-center text-gray-800 text-xs text-gray-400 hover:text-red-600 dark:hover:text-red-600 transition duration-350 ease-in-out"}
        "14 k"]
       [:div {:class "flex-1 flex items-center text-gray-800 dark:text-white text-xs text-gray-400 hover:text-blue-400 dark:hover:text-blue-400 transition duration-350 ease-in-out"}
        ]]]]]
   ]
  )


(defn render-post
  [id]
 (ui/render-ui post-ui "app"))