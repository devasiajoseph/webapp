(ns aupro.feed
  (:require [reagent.core :as r]
            [centipair.ui :as ui]))


(def feed-post (r/atom {:profile-pic "https://pbs.twimg.com/profile_images/1308769664240160770/AfgzWVE7_normal.jpg"
                        :profile-name "Joe Biden"
                        :post-time "Nov 7 2022"
                        :post-intro "America, I’m honored that you have chosen me to lead our great country. The work ahead of us will be hard, but I promise you this: I will be a#Presidentfor all Americans — whether you voted for me or not. I will keep the faith that you have placed in me."
                        :post-pic "https://pbs.twimg.com/media/EnTkhz-XYAEH2kY?format=jpg&name=small"
                        :likes "14k"
                        :comments "9k"
                        :views "100k"}))



(defn feed-post-ui 
  []
  [:div {:class "post-box"}
   
   [:div {:class "flex flex-shrink-0 p-5"}
    [:a {:href (:profile-link feed-post) :class "flex-shrink-0 group block"}
     [:div {:class "flex items-top"}
      [:div
       [:img {:class "inline-block h-10 w-10 rounded-full", :src "https://pbs.twimg.com/profile_images/1308769664240160770/AfgzWVE7_normal.jpg"}]]
      [:div {:class "ml-3"}
       [:p {:class "flex items-center text-base leading-6 font-medium text-gray-800 mt-2"} "Joe Biden\n                "
        [:span {:class "ml-1 text-sm leading-5 font-medium text-gray-400 group-hover:text-gray-300 transition ease-in-out duration-150"} "@JoeBiden . Nov 7 2022"]]]]]]
   [:div {:class "pl-5"}
    [:p {:class "text-base width-auto font-medium text-gray-800 flex-shrink"} "America, I’m honored that you have chosen me to lead our great\n            country.\n            The work ahead of us will be hard, but I promise you this: I\n            will be a"
     [:a {:href "#", :class "text-blue-400 hover:underline"} "#President"] "for all Americans — whether you voted for me or not.\n            I will keep the faith that you have placed in me."]
    [:div {:class "flex my-3 mr-2 "}
     [:img {:class "", :src "https://pbs.twimg.com/media/EnTkhz-XYAEH2kY?format=jpg&name=small"}]]
    [:div {:class "flex"}
     [:div {:class "flex flex-1 items-center text-gray-800 text-xs text-gray-400 hover:text-blue-400 dark:hover:text-blue-400 transition duration-350 ease-in-out"}
      [:img {:src "/static/icons/heart.svg" :class "h-6 w-6"}] "12.3 k"]
     [:div {:class "flex flex-1 items-center text-gray-800 text-xs text-gray-400 hover:text-green-400 dark:hover:text-green-400 transition duration-350 ease-in-out"}
      [:img {:src "/static/icons/comment.svg" :class "h-6 w-6"}] "14 k"]
     [:div {:class "flex flex-1 items-center text-gray-800 text-xs text-gray-400 hover:text-red-600 dark:hover:text-red-600 transition duration-350 ease-in-out"}
      [:img {:src "/static/icons/chart.svg" :class "h-6 w-6"}] "100 k"]
     [:div {:class "flex flex-1 items-center text-gray-800 text-xs text-gray-400 hover:text-red-600 dark:hover:text-red-600 transition duration-350 ease-in-out"}
      [:img {:src "/static/icons/coin.svg" :class "h-6 w-6"}] "25 k"]
     ]
    ]
   ]
  )


(defn render-post
  [id]
 (ui/render-ui feed-post-ui "app"))