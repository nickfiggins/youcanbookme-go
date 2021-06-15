package util

type BookingTimes struct {
	Bookingpaddingminutes int `json:"bookingPaddingMinutes"`
		Breakend              struct {
		} `json:"breakEnd"`
		Breakstart struct {
		} `json:"breakStart"`
		Defaultnumberofslots int    `json:"defaultNumberOfSlots"`
		Fixedend             string `json:"fixedEnd"`
		Fixedstart           string `json:"fixedStart"`
		Friactive            bool   `json:"friActive"`
		Friend               struct {
		} `json:"friEnd"`
		Fristart struct {
		} `json:"friStart"`
		Maxnoticedays    int  `json:"maxNoticeDays"`
		Maxnumberofslots int  `json:"maxNumberOfSlots"`
		Minnoticeminutes int  `json:"minNoticeMinutes"`
		Minnumberofslots int  `json:"minNumberOfSlots"`
		Monactive        bool `json:"monActive"`
		Monend           struct {
		} `json:"monEnd"`
		Monstart struct {
		} `json:"monStart"`
		Ondutytoken string `json:"onDutyToken"`
		Satactive   bool   `json:"satActive"`
		Satend      struct {
		} `json:"satEnd"`
		Satstart struct {
		} `json:"satStart"`
		Slotlengthminutes int    `json:"slotLengthMinutes"`
		Startingdayofweek string `json:"startingDayOfWeek"`
		Sunactive         bool   `json:"sunActive"`
		Sunend            struct {
		} `json:"sunEnd"`
		Sunstart struct {
		} `json:"sunStart"`
		Thuactive bool `json:"thuActive"`
		Thuend    struct {
		} `json:"thuEnd"`
		Thustart struct {
		} `json:"thuStart"`
		Tueactive bool `json:"tueActive"`
		Tueend    struct {
		} `json:"tueEnd"`
		Tuestart struct {
		} `json:"tueStart"`
		Unitsperslot int  `json:"unitsPerSlot"`
		Wedactive    bool `json:"wedActive"`
		Wedend       struct {
		} `json:"wedEnd"`
		Wedstart struct {
		} `json:"wedStart"`
}