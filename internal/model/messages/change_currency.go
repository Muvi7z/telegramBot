package messages

import "context"

func (s *Model) setCurrency(ctx context.Context, msg Message) (text string, err error) {

	return "", err
}

func (s *Model) changeDefaultCurrency() (text string, buttons []map[string]string) {
	btns := make(map[string]string)

	btns["default"] = "stem"

	return "", []map[string]string{btns}
}
