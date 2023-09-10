package task

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/fatih/structs"
)

type TaskYaml struct {
	Tasks []*Task `yaml:"tasks,omitempty" db:"tasks"`
}

type Task struct {
	ID                  int            `yaml:"id,omitempty" db:"id"`                                       // int(11) unsigned NOT NULL DEFAULT 0,
	Type                int            `yaml:"type,omitempty" db:"type"`                                   // tinyint(4) NOT NULL DEFAULT 0,
	Duration            int            `yaml:"duration,omitempty" db:"duration"`                           // int(11) unsigned NOT NULL DEFAULT 0,
	DurationCode        int            `yaml:"duration_code,omitempty" db:"duration_code"`                 // tinyint(4) NOT NULL DEFAULT 0,
	Title               string         `yaml:"title,omitempty" db:"title"`                                 // varchar(100) NOT NULL DEFAULT '',
	Description         string         `yaml:"description,omitempty" db:"description"`                     // text NOT NULL,
	RewardText          string         `yaml:"reward_text,omitempty" db:"reward_text"`                     // varchar(64) NOT NULL DEFAULT '',
	RewardIDList        sql.NullString `yaml:"reward_id_list,omitempty" db:"reward_id_list"`               // text DEFAULT NULL,
	CashReward          int            `yaml:"cash_reward,omitempty" db:"cash_reward"`                     // int(11) unsigned NOT NULL DEFAULT 0,
	ExpReward           int            `yaml:"exp_reward,omitempty" db:"exp_reward"`                       // int(10) NOT NULL DEFAULT 0,
	RewardMethod        int            `yaml:"reward_method,omitempty" db:"reward_method"`                 // tinyint(3) unsigned NOT NULL DEFAULT 0,
	RewardPoints        int            `yaml:"reward_points,omitempty" db:"reward_points"`                 // int(11) NOT NULL DEFAULT 0,
	RewardPointType     int            `yaml:"reward_point_type,omitempty" db:"reward_point_type"`         // int(11) NOT NULL DEFAULT 0,
	MinLevel            int            `yaml:"min_level,omitempty" db:"min_level"`                         // tinyint(3) unsigned NOT NULL DEFAULT 0,
	MaxLevel            int            `yaml:"max_level,omitempty" db:"max_level"`                         // tinyint(3) unsigned NOT NULL DEFAULT 0,
	LevelSpread         int            `yaml:"level_spread,omitempty" db:"level_spread"`                   // int(10) unsigned NOT NULL DEFAULT 0,
	MinPlayers          int            `yaml:"min_players,omitempty" db:"min_players"`                     // int(10) unsigned NOT NULL DEFAULT 0,
	MaxPlayers          int            `yaml:"max_players,omitempty" db:"max_players"`                     // int(10) unsigned NOT NULL DEFAULT 0,
	Repeatable          int            `yaml:"repeatable,omitempty" db:"repeatable"`                       // tinyint(1) unsigned NOT NULL DEFAULT 1,
	FactionReward       int            `yaml:"faction_reward,omitempty" db:"faction_reward"`               // int(10) NOT NULL DEFAULT 0,
	CompletionEmote     string         `yaml:"completion_emote,omitempty" db:"completion_emote"`           // varchar(512) NOT NULL DEFAULT '',
	ReplayTimerGroup    int            `yaml:"replay_timer_group,omitempty" db:"replay_timer_group"`       // int(10) unsigned NOT NULL DEFAULT 0,
	ReplayTimerSeconds  int            `yaml:"replay_timer_seconds,omitempty" db:"replay_timer_seconds"`   // int(10) unsigned NOT NULL DEFAULT 0,
	RequestTimerGroup   int            `yaml:"request_timer_group,omitempty" db:"request_timer_group"`     // int(10) unsigned NOT NULL DEFAULT 0,
	RequestTimerSeconds int            `yaml:"request_timer_seconds,omitempty" db:"request_timer_seconds"` // int(10) unsigned NOT NULL DEFAULT 0,
	DzTemplateId        int            `yaml:"dz_template_id,omitempty" db:"dz_template_id"`               // int(10) unsigned NOT NULL DEFAULT 0,
	LockActivityId      int            `yaml:"lock_activity_id,omitempty" db:"lock_activity_id"`           // int(11) NOT NULL DEFAULT -1,
	FactionAmount       int            `yaml:"faction_amount,omitempty" db:"faction_amount"`               // int(10) NOT NULL DEFAULT 0,
	Activities          []*Activity    `yaml:"activities,omitempty"`
}

type Activity struct {
	TaskID              int            `yaml:"taskid,omitempty" db:"taskid"`                             // int(11) unsigned NOT NULL DEFAULT 0,
	ActivityID          int            `yaml:"activityid" db:"activityid"`                               // int(11) unsigned NOT NULL DEFAULT 0,
	ReqActivityID       int            `yaml:"req_activity_id,omitempty" db:"req_activity_id"`           // int(11) NOT NULL DEFAULT -1,
	Step                int            `yaml:"step,omitempty" db:"step"`                                 // int(11) NOT NULL DEFAULT 0,
	Activitytype        int            `yaml:"activitytype,omitempty" db:"activitytype"`                 // tinyint(3) unsigned NOT NULL DEFAULT 0,
	TargetName          string         `yaml:"target_name,omitempty" db:"target_name"`                   // varchar(64) NOT NULL DEFAULT '',
	Goalmethod          int            `yaml:"goalmethod,omitempty" db:"goalmethod"`                     // int(10) unsigned NOT NULL DEFAULT 0,
	Goalcount           int            `yaml:"goalcount,omitempty" db:"goalcount"`                       // int(11) DEFAULT 1,
	DescriptionOverride string         `yaml:"description_override,omitempty" db:"description_override"` // varchar(128) NOT NULL DEFAULT '',
	NpcMatchList        sql.NullString `yaml:"npc_match_list,omitempty" db:"npc_match_list"`             // text DEFAULT NULL,
	ItemIDList          sql.NullString `yaml:"item_id_list,omitempty" db:"item_id_list"`                 // text DEFAULT NULL,
	ItemList            string         `yaml:"item_list,omitempty" db:"item_list"`                       // varchar(128) NOT NULL DEFAULT '',
	DzSwitchID          int            `yaml:"dz_switch_id,omitempty" db:"dz_switch_id"`                 // int(11) NOT NULL DEFAULT 0,
	MinX                float32        `yaml:"min_x,omitempty" db:"min_x"`                               // float NOT NULL DEFAULT 0,
	MinY                float32        `yaml:"min_y,omitempty" db:"min_y"`                               // float NOT NULL DEFAULT 0,
	MinZ                float32        `yaml:"min_z,omitempty" db:"min_z"`                               // float NOT NULL DEFAULT 0,
	MaxX                float32        `yaml:"max_x,omitempty" db:"max_x"`                               // float NOT NULL DEFAULT 0,
	MaxY                float32        `yaml:"max_y,omitempty" db:"max_y"`                               // float NOT NULL DEFAULT 0,
	MaxZ                float32        `yaml:"max_z,omitempty" db:"max_z"`                               // float NOT NULL DEFAULT 0,
	SkillList           string         `yaml:"skill_list,omitempty" db:"skill_list"`                     // varchar(64) NOT NULL DEFAULT '-1',
	SpellList           string         `yaml:"spell_list,omitempty" db:"spell_list"`                     // varchar(64) NOT NULL DEFAULT '0',
	Zones               string         `yaml:"zones,omitempty" db:"zones"`                               // varchar(64) NOT NULL DEFAULT '',
	ZoneVersion         int            `yaml:"zone_version,omitempty" db:"zone_version"`                 // int(11) DEFAULT -1,
	Optional            int            `yaml:"optional,omitempty" db:"optional"`                         // tinyint(1) NOT NULL DEFAULT 0,
}

func (e *TaskYaml) sanitize() error {
	for _, task := range e.Tasks {
		err := task.sanitize()
		if err != nil {
			return err
		}
	}
	return nil
}

func (task *Task) sanitize() error {
	if task.ID < 1 {
		return fmt.Errorf("task id must be greater than 0")
	}

	for _, activity := range task.Activities {
		err := activity.sanitize()
		if err != nil {
			return err
		}
	}

	return nil
}

func (activity *Activity) sanitize() error {
	if activity.ReqActivityID == -1 {
		activity.ReqActivityID = 0
	}
	return nil
}

func (task *Task) omitEmpty() error {
	baseTask := Task{
		ID: 1,
	}
	err := baseTask.sanitize()
	if err != nil {
		return err
	}
	fields := structs.Fields(task)
	baseFields := structs.Fields(baseTask)

	for fieldIndex, field := range fields {
		if !field.IsExported() {
			continue
		}
		switch field.Kind() {
		case reflect.Int:
			baseVal := baseFields[fieldIndex].Value().(int)
			newVal := field.Value().(int)

			if newVal != baseVal {
				continue
			}
			if newVal == 0 {
				continue
			}
			field.Set(0)
		}
	}
	return nil
}
