package generator

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ECUST-XX/xml"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Generator struct {
	log *logrus.Logger
}

func NewGenerator() *Generator {
	return &Generator{
		log: logrus.New(),
	}
}

func (g *Generator) GenerateModel() ([]byte, error) {
	var matrixHardCoded = [][]int64{
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 0},
	}

	params := g.getParamsFromMatrix(matrixHardCoded)
	if len(params) == 0 {
		return nil, fmt.Errorf("no parameters generated")
	}

	relations := g.getRelations()
	if len(relations) == 0 {
		return nil, fmt.Errorf("no relations generated")
	}

	relationID := relations[0].ID

	model := Model{
		FormatVer: "2.0",
		ID:        uuid.NewString(),
		ShortName: "Model 1",
		Desc:      "Model 1",
		Class: Class{
			ID:        uuid.NewString(),
			ShortName: "Model 1",
			Parameters: Parameters{
				Parameters: params,
			},
			Rules: Rules{
				Rules: g.getRulesFromParams(params, relationID),
			},
		},
		Relations: Relations{
			Relations: relations,
		},
	}

	output, err := xml.MarshalIndentShortForm(model, "", "  ")
	if err != nil {
		g.log.Error("Error generating XML:", err)
		return nil, err
	}

	return output, nil
}

func (g *Generator) getParamsFromMatrix(matrix [][]int64) []Parameter {
	if len(matrix) == 0 {
		g.log.Warn("Matrix is empty")
		return []Parameter{}
	}

	params := make([]Parameter, 0, len(matrix)*len(matrix[0]))
	for x, row := range matrix {
		for y, value := range row {
			if value == 1 {
				continue
			}

			params = append(params, Parameter{
				Type:      "double",
				ID:        uuid.NewString(),
				ShortName: fmt.Sprintf("%d,%d", x, y),
			})
		}
	}

	return params
}

func (g *Generator) getRelations() []Relation {
	relations := make([]Relation, 0, 1)
	relations = append(relations, Relation{
		ID:           uuid.NewString(),
		ShortName:    "y=x+1",
		InObj:        "x:double",
		OutObj:       "y:double",
		RelationType: "simple",
		Content:      "y=x+1",
	})

	return relations
}

func (g *Generator) getRulesFromParams(params []Parameter, relationID string) []Rule {
	if len(params) == 0 {
		g.log.Warn("Empty Params")
		return []Rule{}
	}

	rules := make([]Rule, 0)
	for i, src := range params {
		for j, dst := range params {
			srcVals := g.convertStringsToInts(strings.Split(src.ShortName, ","))
			dstVals := g.convertStringsToInts(strings.Split(dst.ShortName, ","))

			if len(srcVals) != len(dstVals) || len(srcVals) != 2 {
				g.log.Warn(fmt.Sprintf("Invalid params: wrong len after split of %s and %s", src.ShortName, dst.ShortName))
				continue
			}

			xDistance := math.Abs(float64(srcVals[0] - dstVals[0]))
			yDistance := math.Abs(float64(srcVals[1] - dstVals[1]))

			if i == j || xDistance > 1 || yDistance > 1 || (xDistance == 1 && yDistance == 1) {
				continue
			}

			rules = append(rules, Rule{
				ID:        uuid.NewString(),
				ShortName: fmt.Sprintf("%s to %s", src.ShortName, dst.ShortName),
				Relation:  relationID,
				ResultID:  fmt.Sprintf("y:%s", dst.ID),
				InitID:    fmt.Sprintf("x:%s", src.ID),
			})

		}
	}
	return rules
}

func (g *Generator) convertStringsToInts(strs []string) []int64 {
	res := make([]int64, 0, len(strs))

	for _, v := range strs {
		vInt, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			g.log.Error("Error parsing int:", err)
			continue
		}
		res = append(res, vInt)
	}
	return res
}
