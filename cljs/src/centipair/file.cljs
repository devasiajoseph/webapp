(ns centipair.file
  (:require [ajax.core :refer [GET POST]]
            [goog.dom :as gdom]))


(defn upload-file
  [id url]
  (let [f (gdom/getElement id)
        form-data (new js/FormData)
        xhr (new js/XMLHttpRequest)]
    (.append form-data "file" (aget (.-files f) 0))
    (println form-data)))

(defn file-input [field]
  ^{:key (:id @field)}
  [:div {:class "form-control"}
   [:input {:type "file", :class "mx-auto file-input w-full max-w-xs" :id (:id @field)}]
   [:button {:class "btn" :on-click (partial upload-file (:id @field))} "Upload"]
   [:p {:id (str (:id @field) "-message") :class "link-error"} (:message @field)]])