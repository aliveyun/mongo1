package mongo

import (
	//"context"
	db "go.mongodb.org/mongo-driver/mongo"
)

//Client MongoDB client
type Client struct {

	 dbName string
	 uri    string
	 store *MgoStore
	 Con  * db.Client
}

func NewDB(cfg *MgoConf) *Client {

	dbCli, err := NewClient(cfg)
	if err != nil {
		panic(err)
	}
	dbStore := NewStore(dbCli, cfg.DB)

	return &Client{
		Con:dbCli,
		store:dbStore,
	}
}
func (c *Client) Store() *MgoStore {
	
	return c.store
}


/*func (c *Client) InsertOne(ctx context.Context, col mongo.Collection) error {
	err := c.store.InsertOne(context.TODO(), col)
	return err
}

func (c *Client) NewOneFinder(col mongo.Collection) *mongo.OneFinder {
	return mongo.NewOneFinder(col)
}

func (c *Client) FindOne(ctx context.Context, o *mongo.OneFinder) (bool, error) {
	return c.store.FindOne(ctx, o) 
}
func (c *Client)NewFinder(col mongo.Collection) *mongo.Finder {
	return mongo.NewFinder(col)
	
}

func (c *Client) FindMany(ctx context.Context, o *mongo.Finder) error {
	return c.store.FindMany(ctx ,o)
}*/