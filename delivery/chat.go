package delivery

// import (
// 	"context"
// 	"encoding/json"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// 	"gitlab.com/hsf-cloud/lib/errors"
// 	"gitlab.com/hsf-cloud/lib/pagination"
// 	libtime "gitlab.com/hsf-cloud/lib/time"
// 	"gitlab.com/hsf-cloud/lib/uid/xid"
// 	"gitlab.com/hsf-cloud/lib/web"

// 	"gitlab.com/hsf-cloud/proto/pkg/e-commerce/chat_service"
// )

// type createChatRoomRequest struct {
// 	StoreID string `json:"store_id"`
// 	Creator string `json:"creator"`
// }

// type createChatRoomResponse struct {
// 	ChatRoomID string   `json:"room_id"`
// 	Users      []string `json:"users"`
// 	Creator    string   `json:"creator"`
// 	StoreID    string   `json:"store_id"`
// }

// func (h *handler) createChatRoom(ctx echo.Context) error {

// 	req := &createChatRoomRequest{}
// 	err := ctx.Bind(req)
// 	if err != nil {
// 		return errors.NewError(errors.ErrorInvalidInput, err.Error())
// 	}

// 	resp, err := h.chatSVC.ListChatRoom(ctx.Request().Context(), &chat_service.ListChatRoomRequest{
// 		CustomIdentifier: req.StoreID,
// 		UserID:           req.Creator,
// 		Perpage:          2,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	for _, r := range resp.Data {
// 		if r.Creator == req.Creator {
// 			out := &createChatRoomResponse{
// 				ChatRoomID: r.RoomID,
// 				Creator:    r.Creator,
// 				StoreID:    r.CustomIdentifier,
// 				Users:      r.Users,
// 			}
// 			ctx.JSON(http.StatusCreated, out)
// 			return nil
// 		}
// 	}

// 	roomID := xid.NewXIDGenerator().GenUID()
// 	_, err = h.chatSVC.CreateChatRoom(ctx.Request().Context(), &chat_service.CreateChatRoomRequest{
// 		CustomIdentifier: req.StoreID,
// 		Creator:          req.Creator,
// 		AutoReply:        true,
// 		//todo: append user
// 		Users: []string{req.Creator},
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	resp, err = h.chatSVC.ListChatRoom(ctx.Request().Context(), &chat_service.ListChatRoomRequest{
// 		CustomIdentifier: req.StoreID,
// 		UserID:           req.Creator,
// 		Perpage:          1,
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	var users []string
// 	for _, v := range resp.Data {
// 		users = v.Users
// 	}

// 	out := &createChatRoomResponse{
// 		ChatRoomID: roomID,
// 		Users:      users,
// 		StoreID:    req.StoreID,
// 		Creator:    req.Creator,
// 	}

// 	return ctx.JSON(http.StatusCreated, out)

// }

// type createMessageRequest struct {
// 	RoomID string `param:"roomID"`
// 	UserID string `json:"user_id"`
// 	//TBD
// 	UserRole string          `json:"user_role"`
// 	Content  string          `json:"content"`
// 	Meta     json.RawMessage `json:"meta"`
// }

// func (h handler) createMessage(ctx echo.Context) error {

// 	req := &createMessageRequest{}
// 	err := ctx.Bind(req)
// 	if err != nil {
// 		return errors.NewError(errors.ErrorInvalidInput, err.Error())
// 	}

// 	_, err = h.chatSVC.CreateChatMessage(ctx.Request().Context(), &chat_service.CreateChatMessageRequest{
// 		RoomID:    req.RoomID,
// 		UserID:    req.UserID,
// 		UserRole:  chat_service.Role(chat_service.Role_value[req.UserRole]),
// 		Content:   req.Content,
// 		Meta:      req.Meta,
// 		Timestamp: libtime.NowMS(),
// 	})

// 	if err != nil {
// 		return err
// 	}

// 	return ctx.NoContent(http.StatusNoContent)
// }

// type listChatMessagesRequest struct {
// 	RoomID    string `param:"roomID"`
// 	Timestamp int64  `query:"timestamp"`
// 	Page      uint32 `query:"page"`
// 	PerPage   uint32 `query:"perPage"`
// }

// type listChatMessagesResponse struct {
// 	web.ResponsePayLoadMetaData

// 	Data []*StoreChatMessage `json:"data"`
// }

// type StoreChatMessage struct {
// 	Role      string          `json:"role"`
// 	UserID    string          `json:"user_id"`
// 	RoomID    string          `json:"room_id"`
// 	Category  int64           `json:"category"`
// 	Content   string          `json:"content"`
// 	MetaData  json.RawMessage `json:"meta"`
// 	CreatedAt int64           `json:"created_at"`
// }

// func (h handler) listChatMessages(ctx echo.Context) error {

// 	req := &listChatMessagesRequest{}
// 	err := ctx.Bind(req)
// 	if err != nil {
// 		return errors.NewError(errors.ErrorInvalidInput, err.Error())
// 	}

// 	resp, err := h.chatSVC.ListChatMessage(context.Background(), &chat_service.ListChatMessageRequest{
// 		RoomID:       req.RoomID,
// 		CreatedAtGte: req.Timestamp,
// 		Page:         int32(req.Page),
// 		Perpage:      int32(req.PerPage),
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	var list []*StoreChatMessage

// 	for _, v := range resp.Data {
// 		list = append(list, &StoreChatMessage{
// 			Role:      v.Role.String(),
// 			UserID:    v.UserID,
// 			RoomID:    v.RoomID,
// 			Category:  v.Category,
// 			Content:   v.Content,
// 			MetaData:  v.Meta,
// 			CreatedAt: v.CreatedAt,
// 		})
// 	}

// 	out := &listChatMessagesResponse{
// 		ResponsePayLoadMetaData: web.ResponsePayLoadMetaData{
// 			Pagination: &pagination.Pagination{
// 				TotalCount: uint32(resp.Pagination.TotalCount),
// 				TotalPage:  uint32(resp.Pagination.TotalPage),
// 				PerPage:    uint32(resp.Pagination.Perpage),
// 				Page:       uint32(resp.Pagination.Page),
// 			},
// 		},
// 		Data: list,
// 	}

// 	return ctx.JSON(http.StatusOK, out)
// }

// type createStoreAIBotRequest struct {
// 	StoreID             string  `param:"storeID" json:"store_id"`
// 	SystemPrompt        string  `json:"system_prompt"`
// 	EmbeddingConfidence float64 `json:"embedding_confidence"`
// 	TopP                float64 `json:"top_p"`
// 	Temperature         float64 `json:"temperature"`
// 	PresencePenalty     float64 `json:"presence_penalty"`
// 	FrequencyPenalty    float64 `json:"frequency_penalty"`
// }

// func (h handler) createStoreAIBot(ctx echo.Context) error {

// 	req := &createStoreAIBotRequest{}
// 	err := ctx.Bind(req)
// 	if err != nil {
// 		return errors.NewError(errors.ErrorInvalidInput, err.Error())
// 	}
// 	req.StoreID = ctx.Param("storeID")

// 	_, err = h.chatAI.CreateAIChatBot(ctx.Request().Context(), &chat_service.CreateAIChatBotRequest{
// 		CustomIdentifier:    req.StoreID,
// 		SystemPrompt:        req.SystemPrompt,
// 		EmbeddingConfidence: req.EmbeddingConfidence,
// 		TopP:                req.TopP,
// 		Temperature:         req.Temperature,
// 		PresencePenalty:     req.PresencePenalty,
// 		FrequencyPenalty:    req.FrequencyPenalty,
// 		//todo: from token
// 		Creator: xid.NewXIDGenerator().GenUID(),
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	return ctx.NoContent(http.StatusNoContent)
// }

// type createEmbeddingRequest struct {
// 	StoreID       string `param:"store_id" json:"store_id"`
// 	Category      string `json:"category"`
// 	OriginContent string `json:"origin_content"`
// }

// func (h handler) createStoreAIEmbedding(ctx echo.Context) error {
// 	req := &createEmbeddingRequest{}
// 	err := ctx.Bind(req)
// 	if err != nil {
// 		return errors.NewError(errors.ErrorInvalidInput, err.Error())
// 	}
// 	req.StoreID = ctx.Param("storeID")

// 	_, err = h.chatAI.CreateEmbedding(ctx.Request().Context(), &chat_service.CreateEmbeddingRequest{
// 		CustomIdentifier: req.StoreID,
// 		OriginContent:    req.OriginContent,
// 		Category:         req.Category,

// 		//from token
// 		Creator: xid.NewXIDGenerator().GenUID(),
// 	})

// 	if err != nil {
// 		return err
// 	}

// 	return ctx.NoContent(http.StatusNoContent)
// }
