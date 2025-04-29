package variables

type Page struct {
	Title string         `json:"title"`
	User  *User          `json:"user"`
	Data  map[string]any `json:"data"`
}

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Topic   *Topic `json:"topic"`
	Author  *User  `json:"author"`
	Date    string `json:"date"`
}

type Topic struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type User struct {
	ID        []byte `json:"id"`
	Email     string `json:"email"`
	UserName  string `json:"username"`
	LastName  string `json:"lastname"`
	FirstName string `json:"firstname"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
	PassWord  []byte `json:"password"`
}

type Error struct {
	Error string `json:"error"`
}
