(ns centipair.core
  (:require [app.router :as app]
            [centipair.components.notifier :as notifier]
            [centipair.control :as control]))


;;(app/render-home)
(app/init!)
(notifier/render-notifier-component)
(control/load-auth)