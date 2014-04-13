package main

import (
	"view"
	"view/theme"
	"view/layout"
	"view/widget"
	"view/event"
	"fmt"
)

 
func main() {
	var waitOnExit chan bool
	fmt.Println(theme.HexRGBA(0xff77ff00))
	win := view.NewWindow("Mouse Integration Test", 0, 0, 800, 600)
	win.SetName("Test App")
	win.SetSize(600, 400)
	l := layout.NewAbsolute(win)
	tb := widget.NewTextBox(win, msg)
	style := tb.Style()
	style.SetOverflowY(view.STYLE_OVERFLOW_Y_SCROLL)
	
	tb.AddMouseEnterHandler(func (mp event.Mouse){
//		fmt.Println("Mouse Enter:", mp.X, mp.Y)
	})
	
	tb.AddMouseExitHandler(func (mp event.Mouse){
//		fmt.Println("Mouse Exit:", mp.X, mp.Y)
	})
	
	tb.AddMousePositionHandler(func (mp event.Mouse){
//		fmt.Println("Mouse Position:", mp.X, mp.Y)
	})
	
	tb.AddMouseButtonPressHandler(func (mp event.Mouse){
		switch mp.Button {
			case event.MOUSE_BUTTON_LEFT:
				fmt.Println("LEFT PRESS")
			case event.MOUSE_BUTTON_MIDDLE:
				fmt.Println("MIDDLE PRESS")
			case event.MOUSE_BUTTON_RIGHT:
				fmt.Println("RIGHT PRESS")
		}
	})
	
	tb.AddMouseButtonReleaseHandler(func (mp event.Mouse){
		switch mp.Button {
			case event.MOUSE_BUTTON_LEFT:
				fmt.Println("LEFT RELEASE")
			case event.MOUSE_BUTTON_MIDDLE:
				fmt.Println("MIDDLE RELEASE")
			case event.MOUSE_BUTTON_RIGHT:
				fmt.Println("RIGHT RELEASE")
		}
	})
	
	tb.AddMouseWheelUpHandler(func (mp event.Mouse) {
		fmt.Println("UP")
	})
	
	tb.AddMouseWheelDownHandler(func (mp event.Mouse) {
		fmt.Println("DOWN")
	})
	
	tb.AddFocusGainedHandler(func() {
		fmt.Println("Focus Gained")
	})
	
	tb.AddFocusLostHandler(func() {
		fmt.Println("Focus Lost")
	})
	
	l.Add(tb, view.Bounds{0,0,view.Size{300,400}}) 
	win.SetLayout(l)
	win.Start()
	<-waitOnExit
}

const msg = `ABCDEFGHIJKLMNOPQRSTUVWXYZ
abcdefghijklmnopqrstuvwxyz
Banh mi letterpress cred, asymmetrical Williamsburg brunch fixie blog. Direct trade Schlitz Helvetica stumptown chambray, sriracha distillery slow-carb. Ethnic Etsy fanny pack gastropub PBR, skateboard put a bird on it artisan mumblecore Thundercats meggings messenger bag umami chillwave single-origin coffee. Cosby sweater Echo Park tote bag small batch chambray Truffaut Schlitz, bespoke Williamsburg hoodie fingerstache street art tousled cray deep v. Ethnic ennui farm-to-table, fap jean shorts salvia normcore fixie chillwave cred twee freegan DIY Banksy. Forage lo-fi kale chips plaid, YOLO vegan aesthetic 3 wolf moon mixtape deep v Pinterest High Life hoodie. Meh cliche wolf, messenger bag shabby chic craft beer mustache try-hard ethical hoodie typewriter blog organic Bushwick food truck.

Cred irony sartorial aesthetic fingerstache mixtape, ethnic Odd Future. Small batch ethical 3 wolf moon aesthetic farm-to-table jean shorts, bitters slow-carb. Pop-up vegan seitan brunch, sartorial PBR fashion axe Neutra Echo Park VHS. DIY wayfarers Truffaut swag tattooed, keffiyeh farm-to-table seitan yr readymade fingerstache. Wayfarers Neutra Cosby sweater tofu fanny pack, four loko banjo chambray pop-up distillery Wes Anderson artisan 3 wolf moon. Post-ironic bicycle rights authentic, Pitchfork gastropub fixie craft beer next level vinyl paleo Shoreditch PBR stumptown trust fund McSweeney's. High Life ugh narwhal try-hard umami brunch.

Fixie meh literally, four loko dreamcatcher banjo tofu. Cornhole meggings cred, selfies 3 wolf moon trust fund VHS dreamcatcher Schlitz asymmetrical try-hard ugh salvia. Lo-fi iPhone pop-up, twee Tumblr Schlitz fap leggings McSweeney's before they sold out skateboard vegan Cosby sweater raw denim. Slow-carb Portland butcher messenger bag meggings. Carles ugh XOXO before they sold out. Cray whatever roof party Brooklyn, keffiyeh Truffaut food truck small batch 8-bit viral sustainable jean shorts fap Austin before they sold out. Aesthetic hoodie wayfarers, flannel gentrify XOXO plaid dreamcatcher chambray fingerstache food truck Wes Anderson umami Truffaut.
`

