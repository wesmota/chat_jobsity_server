package chatrooms

import "context"

func (suite *ChatsTestSuite) TestListChatRooms() {
	ctx := context.Background()
	chatRooms, err := suite.repo.ListChatRooms(ctx)
	suite.Require().NoError(err)
	suite.Require().NotNil(chatRooms)
	suite.Require().Equal(2, len(chatRooms))
}
