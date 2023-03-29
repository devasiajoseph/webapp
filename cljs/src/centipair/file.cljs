(ns centipair.file
  (:require [ajax.core :refer [GET POST]]
            [centipair.dom :as dom]))



(defn update-progress
  "Updates progress bar"
  [field  progress]
  (swap! field assoc :progress progress))

(defn upload-file
  [field]
  (let [f (dom/get-element (:id @field))
        form-data (new js/FormData)
        xhr (new js/XMLHttpRequest)]
    (.append form-data (:id @field) (aget (.-files f) 0))
    (.open xhr "POST" (:url @field))
    (.setRequestHeader xhr "X-CSRF-Token" (dom/get-value "csrf_token"))
    (set! (.-onprogress xhr.upload)
          (fn [evt]
            (let [progress (* 100 (/ (.-loaded evt) (.-total evt)))]
              (update-progress field progress))))
    (set! (.-onreadystatechange xhr)
          (fn []
            (if (and (= (.-readyState xhr) 4) (= (.-status xhr) 200)) 
              (let [response (js->clj (.parse js/JSON (.-responseText xhr)) :keywordize-keys true)] 
                (swap! field assoc :value (:src response))
                (if (:callback @field) ((:callback @field)))))))
    (.send xhr form-data)))

(defn file-input [field]
  ^{:key (:id @field)}
  [:div {:class ""}[:div {:class "form-control"}
                    [:input {:type "file", :class "mx-auto file-input w-full max-w-xs" :id (:id @field)}]
                    [:button {:class "btn mx-auto" :on-click (partial upload-file field)} "Upload"]
                    [:p {:id (str (:id @field) "-message") :class "link-error"} (:message @field)]
                    [:progress {:class "progress progress-primary mt-5", :value (:progress @field), :max "100"}]]])