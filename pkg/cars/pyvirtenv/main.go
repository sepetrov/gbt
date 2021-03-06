package pyvirtenv

import (
    "path"

    "github.com/jtyr/gbt/pkg/core/car"
    "github.com/jtyr/gbt/pkg/core/utils"
)

// Car inherits the core.Car.
type Car struct {
    car.Car
}

// Init initializes the car.
func (c *Car) Init() {
    defaultRootBg := utils.GetEnv("GBT_CAR_BG", "222")
    defaultRootFg := utils.GetEnv("GBT_CAR_FG", "black")
    defaultRootFm := utils.GetEnv("GBT_CAR_FM", "none")

    c.Model = map[string]car.ModelElement {
        "root": {
            Bg: utils.GetEnv("GBT_CAR_PYVIRTENV_BG", defaultRootBg),
            Fg: utils.GetEnv("GBT_CAR_PYVIRTENV_FG", defaultRootFg),
            Fm: utils.GetEnv("GBT_CAR_PYVIRTENV_FM", defaultRootFm),
            Text: utils.GetEnv("GBT_CAR_PYVIRTENV_FORMAT", " {{ Icon }} {{ Name }} "),
        },
        "Icon": {
            Bg: utils.GetEnv(
                "GBT_CAR_PYVIRTENV_ICON_BG", utils.GetEnv(
                    "GBT_CAR_PYVIRTENV_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_PYVIRTENV_ICON_FG", utils.GetEnv(
                    "GBT_CAR_PYVIRTENV_FG", "33")),
            Fm: utils.GetEnv(
                "GBT_CAR_PYVIRTENV_ICON_FM", utils.GetEnv(
                    "GBT_CAR_PYVIRTENV_FM", defaultRootFm)),
            Text: utils.GetEnv("GBT_CAR_PYVIRTENV_NAME_TEXT", "\ue73c"),
        },
        "Name": {
            Bg: utils.GetEnv(
                "GBT_CAR_PYVIRTENV_NAME_BG", utils.GetEnv(
                    "GBT_CAR_PYVIRTENV_BG", defaultRootBg)),
            Fg: utils.GetEnv(
                "GBT_CAR_PYVIRTENV_NAME_FG", utils.GetEnv(
                    "GBT_CAR_PYVIRTENV_FG", defaultRootFg)),
            Fm: utils.GetEnv(
                "GBT_CAR_PYVIRTENV_NAME_FM", utils.GetEnv(
                    "GBT_CAR_PYVIRTENV_FM", defaultRootFm)),
            Text: utils.GetEnv(
                "GBT_CAR_PYVIRTENV_NAME_TEXT",
                path.Base(utils.GetEnv("VIRTUAL_ENV", ""))),
        },
    }

    if utils.GetEnv("VIRTUAL_ENV", "") != "" {
        c.Display = utils.GetEnvBool("GBT_CAR_PYVIRTENV_DISPLAY", true)
    } else {
        c.Display = utils.GetEnvBool("GBT_CAR_PYVIRTENV_DISPLAY", false)
    }

    c.Wrap = utils.GetEnvBool("GBT_CAR_PYVIRTENV_WRAP", false)
    c.Sep = utils.GetEnv("GBT_CAR_PYVIRTENV_SEP", "\000")
}
