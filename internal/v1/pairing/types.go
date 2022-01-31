package pairing

type Message struct {
	ProtocolVersion int         `json:"protocol_version"`
	Payload         interface{} `json:"payload"`
	Type            int         `json:"type"`
	Status          int         `json:"status"`
}

type PariringRequest struct {
	ServiceName string `json:"service_name"`
	ClientName  string `json:"client_name"`
}

type OptionEncoding struct {
	SymbolLength int `json:"symbol_length"`
	Type         int `json:"type"`
}

type OptionsRequest struct {
	OutputEncodings []OptionEncoding `json:"output_encodings"`
	InputEncodings  []OptionEncoding `json:"input_encodings"`
	PreferredRole   int              `json:"preferred_role"`
}

type ConfigurationRequest struct {
	Encoding   OptionEncoding `json:"encoding"`
	ClientRole int            `json:"client_role"`
}

type SecretRequest struct {
	Secret string `json:"secret"`
}
