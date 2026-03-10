package controllers

import (
    "fmt"
    beego "github.com/beego/beego/v2/server/web"
)

type RedirectPartnerController struct {
    beego.Controller
}

func (c *RedirectPartnerController) Get() {
    // Log all tracking params
    fmt.Println("=== Affiliate Click ===")
    fmt.Println("property_id:", c.GetString("property_id"))
    fmt.Println("feed:", c.GetString("feed"))
    fmt.Println("referral_id:", c.GetString("referral_id"))
    fmt.Println("menu_id:", c.GetString("menu_id"))
    fmt.Println("currency:", c.GetString("currency"))
    fmt.Println("user_type:", c.GetString("user_type"))
    fmt.Println("======================")

    // Get the partner URL and redirect
    directRedirectUrl := c.GetString("direct_redirect_url")
    if directRedirectUrl == "" {
        c.Ctx.WriteString("No redirect URL provided")
        return
    }

    c.Redirect(directRedirectUrl, 302)
}