package fyne

// Window describes a user interface window. Depending on the platform an app
// may have many windows or just the one.
type Window interface {
	// Title returns the current window title.
	// This is typically displayed in the window decorations.
	Title() string
	// SetTitle updates the current title of the window.
	SetTitle(string)

	// FullScreen returns whether or not this window is currently full screen.
	FullScreen() bool
	// SetFullScreen changes the requested fullScreen property
	// true for a fullScreen window and false to unset this.
	SetFullScreen(bool)

	// Resize this window to the requested content size.
	// The result may not be exactly as desired due to various desktop or
	// platform constraints.
	Resize(Size)

	// RequestFocus attempts to raise and focus this window.
	// This should only be called when you are sure the user would want this window
	// to steal focus from any current focused window.
	RequestFocus()

	// FixedSize returns whether or not this window should disable resizing.
	FixedSize() bool
	// SetFixedSize sets a hint that states whether the window should be a fixed
	// size or allow resizing.
	SetFixedSize(bool)

	// CenterOnScreen places a window at the center of the monitor
	// the Window object is currently positioned on.
	CenterOnScreen()

	// Padded, normally true, states whether the window should have inner
	// padding so that components do not touch the window edge.
	Padded() bool
	// SetPadded allows applications to specify that a window should have
	// no inner padding. Useful for fullscreen or graphic based applications.
	SetPadded(bool)

	// Icon returns the window icon, this is used in various ways
	// depending on operating system.
	// Most commonly this is displayed on the window border or task switcher.
	Icon() Resource

	// SetIcon sets the icon resource used for this window.
	// If none is set should return the application icon.
	SetIcon(Resource)

	// SetMaster indicates that closing this window should exit the app
	SetMaster()

	// MainMenu gets the content of the window's top level menu.
	MainMenu() *MainMenu

	// SetMainMenu adds a top level menu to this window.
	// The way this is rendered will depend on the loaded driver.
	SetMainMenu(*MainMenu)

	// SetOnClosed sets a function that runs when the window is closed.
	SetOnClosed(func())

	// SetCloseIntercept sets a function that runs instead of closing if defined.
	// [Window.Close] should be called explicitly in the interceptor to close the window.
	//
	// Since: 1.4
	SetCloseIntercept(func())

	// SetOnDropped allows setting a window-wide callback to receive dropped items.
	// The callback function is called with the absolute position of the drop and a
	// slice of all of the dropped URIs.
	//
	// Since 2.4
	SetOnDropped(func(Position, []URI))

	// Show the window on screen.
	Show()
	// Hide the window from the user.
	// This will not destroy the window or cause the app to exit.
	Hide()
	// Close the window.
	// If it is he "master" window the app will Quit.
	// If it is the only open window and no menu is set via [desktop.App]
	// SetSystemTrayMenu the app will also Quit.
	Close()

	// ShowAndRun is a shortcut to show the window and then run the application.
	// This should be called near the end of a main() function as it will block.
	ShowAndRun()

	// Content returns the content of this window.
	Content() CanvasObject
	// SetContent sets the content of this window.
	SetContent(CanvasObject)
	// Canvas returns the canvas context to render in the window.
	// This can be useful to set a key handler for the window, for example.
	Canvas() Canvas

	// Clipboard returns the system clipboard
	//
	// Deprecated: use App.Clipboard() instead.
	Clipboard() Clipboard

	// --- Custom additions below ---

	// GetScreenRect returns the actual screen coordinates and size of the window,
	// including title bar and borders. This can be useful for precise placement
	// or interacting with external window managers.
	// The coordinates are in screen pixels.
	// Returns an error if the operation is not supported or fails on the current platform.
	//
	// Since: <your_desired_version, e.g., 2.X.X-dev>
	GetScreenRect() (Position, Size)

	// GetNativeHandle returns the underlying platform-specific window handle (HWND on Windows, NSWindow* on macOS, XID on Linux).
	// This is for advanced use cases where direct platform API interaction is needed
	// (e.g., embedding external native libraries or custom window manager interactions).
	// The type of the returned handle depends on the OS:
	// - Windows: `uintptr` (an HWND)
	// - macOS: `unsafe.Pointer` (an NSWindow*)
	// - Linux (X11): `uintptr` (an XID)
	// Applications should type-assert this to the appropriate platform-specific type.
	// Returns nil if the driver does not support exposing the native handle or if the window is not yet created.
	//
	// Since: <your_desired_version, e.g., 2.X.X-dev>
	GetNativeHandle() any

	// SetPosition moves the entire window to a new screen coordinate.
	// The position is relative to the top-left corner of the primary screen.
	// This method moves the entire window, including its title bar and borders.
	//
	// Since: <your_desired_version, e.g., 2.X.X-dev>
	SetPosition(pos Position)

	// SetWindowOnPositionChanged sets a function that is called when the window is moved.
	// This includes user drag operations and programmatic position changes.
	// The position reported is the top-left corner of the window's content area (client area).
	// To get the full window position, use GetScreenRect().
	//
	// Since: <your_desired_version, e.g., 2.X.X-dev>
	SetWindowOnPositionChanged(func(pos Position))
}
