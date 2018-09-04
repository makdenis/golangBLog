package documents

type PostDocument struct {
	Id string `bson:"_id"`
	Title string `bson:"title"`
	Content string `bson:"content"`
}
