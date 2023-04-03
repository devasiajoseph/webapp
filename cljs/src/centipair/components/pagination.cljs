(ns centipair.components.pagination)

(defn next-url 
  [list]
  (if (< (* (:page @list) (:limit @list)) (:total @list))
    (str (:url @list) (+ (:page @list) 1))
    (str (:url @list) (:page @list))))

(defn previous-url 
  [list]
  (if (<= (:page @list) 1)
    (str (:url @list) (:page @list))
    (str (:url @list)  (- (:page @list) 1))))

(defn view 
  [list]
  [:div {:class "btn-group"}
   [:a {:class "btn" :href (previous-url list)} "« Previous"]
   [:a {:class "btn" :href (next-url list)} "Next »"]])