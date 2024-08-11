package cloudbeaver

// CreateTeam -
type CreateTeam struct {
	Query         string          `json:"query"`
	OperationName string          `json:"operationName"`
	Variables     CreateVariables `json:"variables"`
}

// CreateVariables -
type CreateVariables struct {
	Description string `json:"description"`
	TeamId      string `json:"teamId"`
	TeamName    string `json:"teamName"`

	IncludeMetaParameters bool `json:"includeMetaParameters"`
	CustomIncludeBase     bool `json:"customIncludeBase"`
}

// CreateResponse -
type CreateResponse struct {
	Data struct {
		Team Team `json:"team"`
	} `json:"data"`
}

// NewCreateTeam -
func NewCreateTeam(teamId, teamName, desc *string) *CreateTeam {
	return &CreateTeam{
		OperationName: "createTeam",
		Query:         "\n    query createTeam($teamId: ID!, $teamName: String, $description: String, $includeMetaParameters: Boolean!) {\n  team: createTeam(\n    teamId: $teamId\n    teamName: $teamName\n    description: $description\n  ) {\n    ...AdminTeamInfo\n  }\n}\n    \n    fragment AdminTeamInfo on AdminTeamInfo {\n  teamId\n  teamName\n  description\n  teamPermissions\n  metaParameters @include(if: $includeMetaParameters)\n}\n    ",
		Variables: CreateVariables{
			Description:           *desc,
			TeamId:                *teamId,
			TeamName:              *teamName,
			IncludeMetaParameters: false,
			CustomIncludeBase:     true,
		},
	}
}

// DeleteTeam -
type DeleteTeam struct {
	OperationName string          `json:"operationName"`
	Query         string          `json:"query"`
	Variables     DeleteVariables `json:"variables"`
}

// DeleteVariables -
type DeleteVariables struct {
	TeamId string `json:"teamId"`
	Force  bool   `json:"force"`
}

// DeleteResponse -
type DeleteResponse struct {
	Data struct {
		DeleteTeam bool `json:"deleteTeam"`
	} `json:"data"`
}

// NewDeleteTeam -
func NewDeleteTeam(teamId *string) *DeleteTeam {
	return &DeleteTeam{
		OperationName: "deleteTeam",
		Query:         "\n    query deleteTeam($teamId: ID!, $force: Boolean) {\n  deleteTeam(teamId: $teamId, force: $force)\n}\n    ",
		Variables: DeleteVariables{
			TeamId: *teamId,
			Force:  true,
		},
	}
}

// GetAllTeam -
type GetAllTeam struct {
	OperationName string          `json:"operationName"`
	Query         string          `json:"query"`
	Variables     GetAllVariables `json:"variables"`
}

// GetAllVariables -
type GetAllVariables struct {
	IncludeMetaParameters bool `json:"includeMetaParameters"`
	CustomIncludeBase     bool `json:"customIncludeBase"`
}

type GetAllResponse struct {
	Data struct {
		Teams []Team `json:"teams"`
	} `json:"data"`
}

// NewGetAllTeams -
func NewGetAllTeams() *GetAllTeam {
	return &GetAllTeam{
		OperationName: "getTeamsList",
		Query:         "\n    query getTeamsList($teamId: ID, $includeMetaParameters: Boolean!) {\n  teams: listTeams(teamId: $teamId) {\n    ...AdminTeamInfo\n  }\n}\n    \n    fragment AdminTeamInfo on AdminTeamInfo {\n  teamId\n  teamName\n  description\n  teamPermissions\n  metaParameters @include(if: $includeMetaParameters)\n}\n    ",
		Variables: GetAllVariables{
			IncludeMetaParameters: false,
			CustomIncludeBase:     true,
		},
	}
}

// GetTeam -
type GetTeam struct {
	OperationName string       `json:"operationName"`
	Query         string       `json:"query"`
	Variables     GetVariables `json:"variables"`
}

// GetVariables -
type GetVariables struct {
	TeamId string `json:"teamId"`

	IncludeMetaParameters bool `json:"includeMetaParameters"`
	CustomIncludeBase     bool `json:"customIncludeBase"`
}

type GetResponse struct {
	Data struct {
		Teams []Team `json:"teams"`
	} `json:"data"`
}

// NewGetTeams -
func NewGetTeams(teamId *string) *GetTeam {
	return &GetTeam{
		OperationName: "getTeamsList",
		Query:         "\n    query getTeamsList($teamId: ID, $includeMetaParameters: Boolean!) {\n  teams: listTeams(teamId: $teamId) {\n    ...AdminTeamInfo\n  }\n}\n    \n    fragment AdminTeamInfo on AdminTeamInfo {\n  teamId\n  teamName\n  description\n  teamPermissions\n  metaParameters @include(if: $includeMetaParameters)\n}\n    ",
		Variables: GetVariables{
			TeamId:                *teamId,
			IncludeMetaParameters: false,
			CustomIncludeBase:     true,
		},
	}
}

// UpdateTeam -
type UpdateTeam struct {
	OperationName string          `json:"operationName"`
	Query         string          `json:"query"`
	Variables     UpdateVariables `json:"variables"`
}

// UpdateVariables -
type UpdateVariables struct {
	TeamId                string `json:"teamId"`
	TeamName              string `json:"teamName"`
	Description           string `json:"description"`
	IncludeMetaParameters bool   `json:"includeMetaParameters"`
	CustomIncludeBase     bool   `json:"customIncludeBase"`
}

type UpdateResponse struct {
	Data struct {
		Team Team `json:"team"`
	}
}

func NewUpdateTeam(teamId, teamName, desc *string) *UpdateTeam {
	return &UpdateTeam{
		OperationName: "updateTeam",
		Query:         "\n    query updateTeam($teamId: ID!, $teamName: String, $description: String, $includeMetaParameters: Boolean!) {\n  team: updateTeam(\n    teamId: $teamId\n    teamName: $teamName\n    description: $description\n  ) {\n    ...AdminTeamInfo\n  }\n}\n    \n    fragment AdminTeamInfo on AdminTeamInfo {\n  teamId\n  teamName\n  description\n  teamPermissions\n  metaParameters @include(if: $includeMetaParameters)\n}\n    ",
		Variables: UpdateVariables{
			TeamId:      *teamId,
			TeamName:    *teamName,
			Description: *desc,

			CustomIncludeBase:     true,
			IncludeMetaParameters: false,
		},
	}
}
