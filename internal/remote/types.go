package remote

type RemoteKeyCode int32

const (
	RemoteKeyCode_KEYCODE_UNKNOWN                       RemoteKeyCode = 0
	RemoteKeyCode_KEYCODE_SOFT_LEFT                     RemoteKeyCode = 1
	RemoteKeyCode_KEYCODE_SOFT_RIGHT                    RemoteKeyCode = 2
	RemoteKeyCode_KEYCODE_HOME                          RemoteKeyCode = 3
	RemoteKeyCode_KEYCODE_BACK                          RemoteKeyCode = 4
	RemoteKeyCode_KEYCODE_CALL                          RemoteKeyCode = 5
	RemoteKeyCode_KEYCODE_ENDCALL                       RemoteKeyCode = 6
	RemoteKeyCode_KEYCODE_0                             RemoteKeyCode = 7
	RemoteKeyCode_KEYCODE_1                             RemoteKeyCode = 8
	RemoteKeyCode_KEYCODE_2                             RemoteKeyCode = 9
	RemoteKeyCode_KEYCODE_3                             RemoteKeyCode = 10
	RemoteKeyCode_KEYCODE_4                             RemoteKeyCode = 11
	RemoteKeyCode_KEYCODE_5                             RemoteKeyCode = 12
	RemoteKeyCode_KEYCODE_6                             RemoteKeyCode = 13
	RemoteKeyCode_KEYCODE_7                             RemoteKeyCode = 14
	RemoteKeyCode_KEYCODE_8                             RemoteKeyCode = 15
	RemoteKeyCode_KEYCODE_9                             RemoteKeyCode = 16
	RemoteKeyCode_KEYCODE_STAR                          RemoteKeyCode = 17
	RemoteKeyCode_KEYCODE_POUND                         RemoteKeyCode = 18
	RemoteKeyCode_KEYCODE_DPAD_UP                       RemoteKeyCode = 19
	RemoteKeyCode_KEYCODE_DPAD_DOWN                     RemoteKeyCode = 20
	RemoteKeyCode_KEYCODE_DPAD_LEFT                     RemoteKeyCode = 21
	RemoteKeyCode_KEYCODE_DPAD_RIGHT                    RemoteKeyCode = 22
	RemoteKeyCode_KEYCODE_DPAD_CENTER                   RemoteKeyCode = 23
	RemoteKeyCode_KEYCODE_VOLUME_UP                     RemoteKeyCode = 24
	RemoteKeyCode_KEYCODE_VOLUME_DOWN                   RemoteKeyCode = 25
	RemoteKeyCode_KEYCODE_POWER                         RemoteKeyCode = 26
	RemoteKeyCode_KEYCODE_CAMERA                        RemoteKeyCode = 27
	RemoteKeyCode_KEYCODE_CLEAR                         RemoteKeyCode = 28
	RemoteKeyCode_KEYCODE_A                             RemoteKeyCode = 29
	RemoteKeyCode_KEYCODE_B                             RemoteKeyCode = 30
	RemoteKeyCode_KEYCODE_C                             RemoteKeyCode = 31
	RemoteKeyCode_KEYCODE_D                             RemoteKeyCode = 32
	RemoteKeyCode_KEYCODE_E                             RemoteKeyCode = 33
	RemoteKeyCode_KEYCODE_F                             RemoteKeyCode = 34
	RemoteKeyCode_KEYCODE_G                             RemoteKeyCode = 35
	RemoteKeyCode_KEYCODE_H                             RemoteKeyCode = 36
	RemoteKeyCode_KEYCODE_I                             RemoteKeyCode = 37
	RemoteKeyCode_KEYCODE_J                             RemoteKeyCode = 38
	RemoteKeyCode_KEYCODE_K                             RemoteKeyCode = 39
	RemoteKeyCode_KEYCODE_L                             RemoteKeyCode = 40
	RemoteKeyCode_KEYCODE_M                             RemoteKeyCode = 41
	RemoteKeyCode_KEYCODE_N                             RemoteKeyCode = 42
	RemoteKeyCode_KEYCODE_O                             RemoteKeyCode = 43
	RemoteKeyCode_KEYCODE_P                             RemoteKeyCode = 44
	RemoteKeyCode_KEYCODE_Q                             RemoteKeyCode = 45
	RemoteKeyCode_KEYCODE_R                             RemoteKeyCode = 46
	RemoteKeyCode_KEYCODE_S                             RemoteKeyCode = 47
	RemoteKeyCode_KEYCODE_T                             RemoteKeyCode = 48
	RemoteKeyCode_KEYCODE_U                             RemoteKeyCode = 49
	RemoteKeyCode_KEYCODE_V                             RemoteKeyCode = 50
	RemoteKeyCode_KEYCODE_W                             RemoteKeyCode = 51
	RemoteKeyCode_KEYCODE_X                             RemoteKeyCode = 52
	RemoteKeyCode_KEYCODE_Y                             RemoteKeyCode = 53
	RemoteKeyCode_KEYCODE_Z                             RemoteKeyCode = 54
	RemoteKeyCode_KEYCODE_COMMA                         RemoteKeyCode = 55
	RemoteKeyCode_KEYCODE_PERIOD                        RemoteKeyCode = 56
	RemoteKeyCode_KEYCODE_ALT_LEFT                      RemoteKeyCode = 57
	RemoteKeyCode_KEYCODE_ALT_RIGHT                     RemoteKeyCode = 58
	RemoteKeyCode_KEYCODE_SHIFT_LEFT                    RemoteKeyCode = 59
	RemoteKeyCode_KEYCODE_SHIFT_RIGHT                   RemoteKeyCode = 60
	RemoteKeyCode_KEYCODE_TAB                           RemoteKeyCode = 61
	RemoteKeyCode_KEYCODE_SPACE                         RemoteKeyCode = 62
	RemoteKeyCode_KEYCODE_SYM                           RemoteKeyCode = 63
	RemoteKeyCode_KEYCODE_EXPLORER                      RemoteKeyCode = 64
	RemoteKeyCode_KEYCODE_ENVELOPE                      RemoteKeyCode = 65
	RemoteKeyCode_KEYCODE_ENTER                         RemoteKeyCode = 66
	RemoteKeyCode_KEYCODE_DEL                           RemoteKeyCode = 67
	RemoteKeyCode_KEYCODE_GRAVE                         RemoteKeyCode = 68
	RemoteKeyCode_KEYCODE_MINUS                         RemoteKeyCode = 69
	RemoteKeyCode_KEYCODE_EQUALS                        RemoteKeyCode = 70
	RemoteKeyCode_KEYCODE_LEFT_BRACKET                  RemoteKeyCode = 71
	RemoteKeyCode_KEYCODE_RIGHT_BRACKET                 RemoteKeyCode = 72
	RemoteKeyCode_KEYCODE_BACKSLASH                     RemoteKeyCode = 73
	RemoteKeyCode_KEYCODE_SEMICOLON                     RemoteKeyCode = 74
	RemoteKeyCode_KEYCODE_APOSTROPHE                    RemoteKeyCode = 75
	RemoteKeyCode_KEYCODE_SLASH                         RemoteKeyCode = 76
	RemoteKeyCode_KEYCODE_AT                            RemoteKeyCode = 77
	RemoteKeyCode_KEYCODE_NUM                           RemoteKeyCode = 78
	RemoteKeyCode_KEYCODE_HEADSETHOOK                   RemoteKeyCode = 79
	RemoteKeyCode_KEYCODE_FOCUS                         RemoteKeyCode = 80
	RemoteKeyCode_KEYCODE_PLUS                          RemoteKeyCode = 81
	RemoteKeyCode_KEYCODE_MENU                          RemoteKeyCode = 82
	RemoteKeyCode_KEYCODE_NOTIFICATION                  RemoteKeyCode = 83
	RemoteKeyCode_KEYCODE_SEARCH                        RemoteKeyCode = 84
	RemoteKeyCode_KEYCODE_MEDIA_PLAY_PAUSE              RemoteKeyCode = 85
	RemoteKeyCode_KEYCODE_MEDIA_STOP                    RemoteKeyCode = 86
	RemoteKeyCode_KEYCODE_MEDIA_NEXT                    RemoteKeyCode = 87
	RemoteKeyCode_KEYCODE_MEDIA_PREVIOUS                RemoteKeyCode = 88
	RemoteKeyCode_KEYCODE_MEDIA_REWIND                  RemoteKeyCode = 89
	RemoteKeyCode_KEYCODE_MEDIA_FAST_FORWARD            RemoteKeyCode = 90
	RemoteKeyCode_KEYCODE_MUTE                          RemoteKeyCode = 91
	RemoteKeyCode_KEYCODE_PAGE_UP                       RemoteKeyCode = 92
	RemoteKeyCode_KEYCODE_PAGE_DOWN                     RemoteKeyCode = 93
	RemoteKeyCode_KEYCODE_PICTSYMBOLS                   RemoteKeyCode = 94
	RemoteKeyCode_KEYCODE_SWITCH_CHARSET                RemoteKeyCode = 95
	RemoteKeyCode_KEYCODE_BUTTON_A                      RemoteKeyCode = 96
	RemoteKeyCode_KEYCODE_BUTTON_B                      RemoteKeyCode = 97
	RemoteKeyCode_KEYCODE_BUTTON_C                      RemoteKeyCode = 98
	RemoteKeyCode_KEYCODE_BUTTON_X                      RemoteKeyCode = 99
	RemoteKeyCode_KEYCODE_BUTTON_Y                      RemoteKeyCode = 100
	RemoteKeyCode_KEYCODE_BUTTON_Z                      RemoteKeyCode = 101
	RemoteKeyCode_KEYCODE_BUTTON_L1                     RemoteKeyCode = 102
	RemoteKeyCode_KEYCODE_BUTTON_R1                     RemoteKeyCode = 103
	RemoteKeyCode_KEYCODE_BUTTON_L2                     RemoteKeyCode = 104
	RemoteKeyCode_KEYCODE_BUTTON_R2                     RemoteKeyCode = 105
	RemoteKeyCode_KEYCODE_BUTTON_THUMBL                 RemoteKeyCode = 106
	RemoteKeyCode_KEYCODE_BUTTON_THUMBR                 RemoteKeyCode = 107
	RemoteKeyCode_KEYCODE_BUTTON_START                  RemoteKeyCode = 108
	RemoteKeyCode_KEYCODE_BUTTON_SELECT                 RemoteKeyCode = 109
	RemoteKeyCode_KEYCODE_BUTTON_MODE                   RemoteKeyCode = 110
	RemoteKeyCode_KEYCODE_ESCAPE                        RemoteKeyCode = 111
	RemoteKeyCode_KEYCODE_FORWARD_DEL                   RemoteKeyCode = 112
	RemoteKeyCode_KEYCODE_CTRL_LEFT                     RemoteKeyCode = 113
	RemoteKeyCode_KEYCODE_CTRL_RIGHT                    RemoteKeyCode = 114
	RemoteKeyCode_KEYCODE_CAPS_LOCK                     RemoteKeyCode = 115
	RemoteKeyCode_KEYCODE_SCROLL_LOCK                   RemoteKeyCode = 116
	RemoteKeyCode_KEYCODE_META_LEFT                     RemoteKeyCode = 117
	RemoteKeyCode_KEYCODE_META_RIGHT                    RemoteKeyCode = 118
	RemoteKeyCode_KEYCODE_FUNCTION                      RemoteKeyCode = 119
	RemoteKeyCode_KEYCODE_SYSRQ                         RemoteKeyCode = 120
	RemoteKeyCode_KEYCODE_BREAK                         RemoteKeyCode = 121
	RemoteKeyCode_KEYCODE_MOVE_HOME                     RemoteKeyCode = 122
	RemoteKeyCode_KEYCODE_MOVE_END                      RemoteKeyCode = 123
	RemoteKeyCode_KEYCODE_INSERT                        RemoteKeyCode = 124
	RemoteKeyCode_KEYCODE_FORWARD                       RemoteKeyCode = 125
	RemoteKeyCode_KEYCODE_MEDIA_PLAY                    RemoteKeyCode = 126
	RemoteKeyCode_KEYCODE_MEDIA_PAUSE                   RemoteKeyCode = 127
	RemoteKeyCode_KEYCODE_MEDIA_CLOSE                   RemoteKeyCode = 128
	RemoteKeyCode_KEYCODE_MEDIA_EJECT                   RemoteKeyCode = 129
	RemoteKeyCode_KEYCODE_MEDIA_RECORD                  RemoteKeyCode = 130
	RemoteKeyCode_KEYCODE_F1                            RemoteKeyCode = 131
	RemoteKeyCode_KEYCODE_F2                            RemoteKeyCode = 132
	RemoteKeyCode_KEYCODE_F3                            RemoteKeyCode = 133
	RemoteKeyCode_KEYCODE_F4                            RemoteKeyCode = 134
	RemoteKeyCode_KEYCODE_F5                            RemoteKeyCode = 135
	RemoteKeyCode_KEYCODE_F6                            RemoteKeyCode = 136
	RemoteKeyCode_KEYCODE_F7                            RemoteKeyCode = 137
	RemoteKeyCode_KEYCODE_F8                            RemoteKeyCode = 138
	RemoteKeyCode_KEYCODE_F9                            RemoteKeyCode = 139
	RemoteKeyCode_KEYCODE_F10                           RemoteKeyCode = 140
	RemoteKeyCode_KEYCODE_F11                           RemoteKeyCode = 141
	RemoteKeyCode_KEYCODE_F12                           RemoteKeyCode = 142
	RemoteKeyCode_KEYCODE_NUM_LOCK                      RemoteKeyCode = 143
	RemoteKeyCode_KEYCODE_NUMPAD_0                      RemoteKeyCode = 144
	RemoteKeyCode_KEYCODE_NUMPAD_1                      RemoteKeyCode = 145
	RemoteKeyCode_KEYCODE_NUMPAD_2                      RemoteKeyCode = 146
	RemoteKeyCode_KEYCODE_NUMPAD_3                      RemoteKeyCode = 147
	RemoteKeyCode_KEYCODE_NUMPAD_4                      RemoteKeyCode = 148
	RemoteKeyCode_KEYCODE_NUMPAD_5                      RemoteKeyCode = 149
	RemoteKeyCode_KEYCODE_NUMPAD_6                      RemoteKeyCode = 150
	RemoteKeyCode_KEYCODE_NUMPAD_7                      RemoteKeyCode = 151
	RemoteKeyCode_KEYCODE_NUMPAD_8                      RemoteKeyCode = 152
	RemoteKeyCode_KEYCODE_NUMPAD_9                      RemoteKeyCode = 153
	RemoteKeyCode_KEYCODE_NUMPAD_DIVIDE                 RemoteKeyCode = 154
	RemoteKeyCode_KEYCODE_NUMPAD_MULTIPLY               RemoteKeyCode = 155
	RemoteKeyCode_KEYCODE_NUMPAD_SUBTRACT               RemoteKeyCode = 156
	RemoteKeyCode_KEYCODE_NUMPAD_ADD                    RemoteKeyCode = 157
	RemoteKeyCode_KEYCODE_NUMPAD_DOT                    RemoteKeyCode = 158
	RemoteKeyCode_KEYCODE_NUMPAD_COMMA                  RemoteKeyCode = 159
	RemoteKeyCode_KEYCODE_NUMPAD_ENTER                  RemoteKeyCode = 160
	RemoteKeyCode_KEYCODE_NUMPAD_EQUALS                 RemoteKeyCode = 161
	RemoteKeyCode_KEYCODE_NUMPAD_LEFT_PAREN             RemoteKeyCode = 162
	RemoteKeyCode_KEYCODE_NUMPAD_RIGHT_PAREN            RemoteKeyCode = 163
	RemoteKeyCode_KEYCODE_VOLUME_MUTE                   RemoteKeyCode = 164
	RemoteKeyCode_KEYCODE_INFO                          RemoteKeyCode = 165
	RemoteKeyCode_KEYCODE_CHANNEL_UP                    RemoteKeyCode = 166
	RemoteKeyCode_KEYCODE_CHANNEL_DOWN                  RemoteKeyCode = 167
	RemoteKeyCode_KEYCODE_ZOOM_IN                       RemoteKeyCode = 168
	RemoteKeyCode_KEYCODE_ZOOM_OUT                      RemoteKeyCode = 169
	RemoteKeyCode_KEYCODE_TV                            RemoteKeyCode = 170
	RemoteKeyCode_KEYCODE_WINDOW                        RemoteKeyCode = 171
	RemoteKeyCode_KEYCODE_GUIDE                         RemoteKeyCode = 172
	RemoteKeyCode_KEYCODE_DVR                           RemoteKeyCode = 173
	RemoteKeyCode_KEYCODE_BOOKMARK                      RemoteKeyCode = 174
	RemoteKeyCode_KEYCODE_CAPTIONS                      RemoteKeyCode = 175
	RemoteKeyCode_KEYCODE_SETTINGS                      RemoteKeyCode = 176
	RemoteKeyCode_KEYCODE_TV_POWER                      RemoteKeyCode = 177
	RemoteKeyCode_KEYCODE_TV_INPUT                      RemoteKeyCode = 178
	RemoteKeyCode_KEYCODE_STB_POWER                     RemoteKeyCode = 179
	RemoteKeyCode_KEYCODE_STB_INPUT                     RemoteKeyCode = 180
	RemoteKeyCode_KEYCODE_AVR_POWER                     RemoteKeyCode = 181
	RemoteKeyCode_KEYCODE_AVR_INPUT                     RemoteKeyCode = 182
	RemoteKeyCode_KEYCODE_PROG_RED                      RemoteKeyCode = 183
	RemoteKeyCode_KEYCODE_PROG_GREEN                    RemoteKeyCode = 184
	RemoteKeyCode_KEYCODE_PROG_YELLOW                   RemoteKeyCode = 185
	RemoteKeyCode_KEYCODE_PROG_BLUE                     RemoteKeyCode = 186
	RemoteKeyCode_KEYCODE_APP_SWITCH                    RemoteKeyCode = 187
	RemoteKeyCode_KEYCODE_BUTTON_1                      RemoteKeyCode = 188
	RemoteKeyCode_KEYCODE_BUTTON_2                      RemoteKeyCode = 189
	RemoteKeyCode_KEYCODE_BUTTON_3                      RemoteKeyCode = 190
	RemoteKeyCode_KEYCODE_BUTTON_4                      RemoteKeyCode = 191
	RemoteKeyCode_KEYCODE_BUTTON_5                      RemoteKeyCode = 192
	RemoteKeyCode_KEYCODE_BUTTON_6                      RemoteKeyCode = 193
	RemoteKeyCode_KEYCODE_BUTTON_7                      RemoteKeyCode = 194
	RemoteKeyCode_KEYCODE_BUTTON_8                      RemoteKeyCode = 195
	RemoteKeyCode_KEYCODE_BUTTON_9                      RemoteKeyCode = 196
	RemoteKeyCode_KEYCODE_BUTTON_10                     RemoteKeyCode = 197
	RemoteKeyCode_KEYCODE_BUTTON_11                     RemoteKeyCode = 198
	RemoteKeyCode_KEYCODE_BUTTON_12                     RemoteKeyCode = 199
	RemoteKeyCode_KEYCODE_BUTTON_13                     RemoteKeyCode = 200
	RemoteKeyCode_KEYCODE_BUTTON_14                     RemoteKeyCode = 201
	RemoteKeyCode_KEYCODE_BUTTON_15                     RemoteKeyCode = 202
	RemoteKeyCode_KEYCODE_BUTTON_16                     RemoteKeyCode = 203
	RemoteKeyCode_KEYCODE_LANGUAGE_SWITCH               RemoteKeyCode = 204
	RemoteKeyCode_KEYCODE_MANNER_MODE                   RemoteKeyCode = 205
	RemoteKeyCode_KEYCODE_3D_MODE                       RemoteKeyCode = 206
	RemoteKeyCode_KEYCODE_CONTACTS                      RemoteKeyCode = 207
	RemoteKeyCode_KEYCODE_CALENDAR                      RemoteKeyCode = 208
	RemoteKeyCode_KEYCODE_MUSIC                         RemoteKeyCode = 209
	RemoteKeyCode_KEYCODE_CALCULATOR                    RemoteKeyCode = 210
	RemoteKeyCode_KEYCODE_ZENKAKU_HANKAKU               RemoteKeyCode = 211
	RemoteKeyCode_KEYCODE_EISU                          RemoteKeyCode = 212
	RemoteKeyCode_KEYCODE_MUHENKAN                      RemoteKeyCode = 213
	RemoteKeyCode_KEYCODE_HENKAN                        RemoteKeyCode = 214
	RemoteKeyCode_KEYCODE_KATAKANA_HIRAGANA             RemoteKeyCode = 215
	RemoteKeyCode_KEYCODE_YEN                           RemoteKeyCode = 216
	RemoteKeyCode_KEYCODE_RO                            RemoteKeyCode = 217
	RemoteKeyCode_KEYCODE_KANA                          RemoteKeyCode = 218
	RemoteKeyCode_KEYCODE_ASSIST                        RemoteKeyCode = 219
	RemoteKeyCode_KEYCODE_BRIGHTNESS_DOWN               RemoteKeyCode = 220
	RemoteKeyCode_KEYCODE_BRIGHTNESS_UP                 RemoteKeyCode = 221
	RemoteKeyCode_KEYCODE_MEDIA_AUDIO_TRACK             RemoteKeyCode = 222
	RemoteKeyCode_KEYCODE_SLEEP                         RemoteKeyCode = 223
	RemoteKeyCode_KEYCODE_WAKEUP                        RemoteKeyCode = 224
	RemoteKeyCode_KEYCODE_PAIRING                       RemoteKeyCode = 225
	RemoteKeyCode_KEYCODE_MEDIA_TOP_MENU                RemoteKeyCode = 226
	RemoteKeyCode_KEYCODE_11                            RemoteKeyCode = 227
	RemoteKeyCode_KEYCODE_12                            RemoteKeyCode = 228
	RemoteKeyCode_KEYCODE_LAST_CHANNEL                  RemoteKeyCode = 229
	RemoteKeyCode_KEYCODE_TV_DATA_SERVICE               RemoteKeyCode = 230
	RemoteKeyCode_KEYCODE_VOICE_ASSIST                  RemoteKeyCode = 231
	RemoteKeyCode_KEYCODE_TV_RADIO_SERVICE              RemoteKeyCode = 232
	RemoteKeyCode_KEYCODE_TV_TELETEXT                   RemoteKeyCode = 233
	RemoteKeyCode_KEYCODE_TV_NUMBER_ENTRY               RemoteKeyCode = 234
	RemoteKeyCode_KEYCODE_TV_TERRESTRIAL_ANALOG         RemoteKeyCode = 235
	RemoteKeyCode_KEYCODE_TV_TERRESTRIAL_DIGITAL        RemoteKeyCode = 236
	RemoteKeyCode_KEYCODE_TV_SATELLITE                  RemoteKeyCode = 237
	RemoteKeyCode_KEYCODE_TV_SATELLITE_BS               RemoteKeyCode = 238
	RemoteKeyCode_KEYCODE_TV_SATELLITE_CS               RemoteKeyCode = 239
	RemoteKeyCode_KEYCODE_TV_SATELLITE_SERVICE          RemoteKeyCode = 240
	RemoteKeyCode_KEYCODE_TV_NETWORK                    RemoteKeyCode = 241
	RemoteKeyCode_KEYCODE_TV_ANTENNA_CABLE              RemoteKeyCode = 242
	RemoteKeyCode_KEYCODE_TV_INPUT_HDMI_1               RemoteKeyCode = 243
	RemoteKeyCode_KEYCODE_TV_INPUT_HDMI_2               RemoteKeyCode = 244
	RemoteKeyCode_KEYCODE_TV_INPUT_HDMI_3               RemoteKeyCode = 245
	RemoteKeyCode_KEYCODE_TV_INPUT_HDMI_4               RemoteKeyCode = 246
	RemoteKeyCode_KEYCODE_TV_INPUT_COMPOSITE_1          RemoteKeyCode = 247
	RemoteKeyCode_KEYCODE_TV_INPUT_COMPOSITE_2          RemoteKeyCode = 248
	RemoteKeyCode_KEYCODE_TV_INPUT_COMPONENT_1          RemoteKeyCode = 249
	RemoteKeyCode_KEYCODE_TV_INPUT_COMPONENT_2          RemoteKeyCode = 250
	RemoteKeyCode_KEYCODE_TV_INPUT_VGA_1                RemoteKeyCode = 251
	RemoteKeyCode_KEYCODE_TV_AUDIO_DESCRIPTION          RemoteKeyCode = 252
	RemoteKeyCode_KEYCODE_TV_AUDIO_DESCRIPTION_MIX_UP   RemoteKeyCode = 253
	RemoteKeyCode_KEYCODE_TV_AUDIO_DESCRIPTION_MIX_DOWN RemoteKeyCode = 254
	RemoteKeyCode_KEYCODE_TV_ZOOM_MODE                  RemoteKeyCode = 255
	RemoteKeyCode_KEYCODE_TV_CONTENTS_MENU              RemoteKeyCode = 256
	RemoteKeyCode_KEYCODE_TV_MEDIA_CONTEXT_MENU         RemoteKeyCode = 257
	RemoteKeyCode_KEYCODE_TV_TIMER_PROGRAMMING          RemoteKeyCode = 258
	RemoteKeyCode_KEYCODE_HELP                          RemoteKeyCode = 259
	RemoteKeyCode_KEYCODE_NAVIGATE_PREVIOUS             RemoteKeyCode = 260
	RemoteKeyCode_KEYCODE_NAVIGATE_NEXT                 RemoteKeyCode = 261
	RemoteKeyCode_KEYCODE_NAVIGATE_IN                   RemoteKeyCode = 262
	RemoteKeyCode_KEYCODE_NAVIGATE_OUT                  RemoteKeyCode = 263
	RemoteKeyCode_KEYCODE_STEM_PRIMARY                  RemoteKeyCode = 264
	RemoteKeyCode_KEYCODE_STEM_1                        RemoteKeyCode = 265
	RemoteKeyCode_KEYCODE_STEM_2                        RemoteKeyCode = 266
	RemoteKeyCode_KEYCODE_STEM_3                        RemoteKeyCode = 267
	RemoteKeyCode_KEYCODE_DPAD_UP_LEFT                  RemoteKeyCode = 268
	RemoteKeyCode_KEYCODE_DPAD_DOWN_LEFT                RemoteKeyCode = 269
	RemoteKeyCode_KEYCODE_DPAD_UP_RIGHT                 RemoteKeyCode = 270
	RemoteKeyCode_KEYCODE_DPAD_DOWN_RIGHT               RemoteKeyCode = 271
	RemoteKeyCode_KEYCODE_MEDIA_SKIP_FORWARD            RemoteKeyCode = 272
	RemoteKeyCode_KEYCODE_MEDIA_SKIP_BACKWARD           RemoteKeyCode = 273
	RemoteKeyCode_KEYCODE_MEDIA_STEP_FORWARD            RemoteKeyCode = 274
	RemoteKeyCode_KEYCODE_MEDIA_STEP_BACKWARD           RemoteKeyCode = 275
	RemoteKeyCode_KEYCODE_SOFT_SLEEP                    RemoteKeyCode = 276
	RemoteKeyCode_KEYCODE_CUT                           RemoteKeyCode = 277
	RemoteKeyCode_KEYCODE_COPY                          RemoteKeyCode = 278
	RemoteKeyCode_KEYCODE_PASTE                         RemoteKeyCode = 279
	RemoteKeyCode_KEYCODE_SYSTEM_NAVIGATION_UP          RemoteKeyCode = 280
	RemoteKeyCode_KEYCODE_SYSTEM_NAVIGATION_DOWN        RemoteKeyCode = 281
	RemoteKeyCode_KEYCODE_SYSTEM_NAVIGATION_LEFT        RemoteKeyCode = 282
	RemoteKeyCode_KEYCODE_SYSTEM_NAVIGATION_RIGHT       RemoteKeyCode = 283
	RemoteKeyCode_KEYCODE_ALL_APPS                      RemoteKeyCode = 284
	RemoteKeyCode_KEYCODE_REFRESH                       RemoteKeyCode = 285
	RemoteKeyCode_KEYCODE_THUMBS_UP                     RemoteKeyCode = 286
	RemoteKeyCode_KEYCODE_THUMBS_DOWN                   RemoteKeyCode = 287
	RemoteKeyCode_KEYCODE_PROFILE_SWITCH                RemoteKeyCode = 288
	RemoteKeyCode_KEYCODE_VIDEO_APP_1                   RemoteKeyCode = 289
	RemoteKeyCode_KEYCODE_VIDEO_APP_2                   RemoteKeyCode = 290
	RemoteKeyCode_KEYCODE_VIDEO_APP_3                   RemoteKeyCode = 291
	RemoteKeyCode_KEYCODE_VIDEO_APP_4                   RemoteKeyCode = 292
	RemoteKeyCode_KEYCODE_VIDEO_APP_5                   RemoteKeyCode = 293
	RemoteKeyCode_KEYCODE_VIDEO_APP_6                   RemoteKeyCode = 294
	RemoteKeyCode_KEYCODE_VIDEO_APP_7                   RemoteKeyCode = 295
	RemoteKeyCode_KEYCODE_VIDEO_APP_8                   RemoteKeyCode = 296
	RemoteKeyCode_KEYCODE_FEATURED_APP_1                RemoteKeyCode = 297
	RemoteKeyCode_KEYCODE_FEATURED_APP_2                RemoteKeyCode = 298
	RemoteKeyCode_KEYCODE_FEATURED_APP_3                RemoteKeyCode = 299
	RemoteKeyCode_KEYCODE_FEATURED_APP_4                RemoteKeyCode = 300
	RemoteKeyCode_KEYCODE_DEMO_APP_1                    RemoteKeyCode = 301
	RemoteKeyCode_KEYCODE_DEMO_APP_2                    RemoteKeyCode = 302
	RemoteKeyCode_KEYCODE_DEMO_APP_3                    RemoteKeyCode = 303
	RemoteKeyCode_KEYCODE_DEMO_APP_4                    RemoteKeyCode = 304
)

var (
	RemoteKeyCode_value = map[string]int32{
		"KEYCODE_UNKNOWN":                       0,
		"KEYCODE_SOFT_LEFT":                     1,
		"KEYCODE_SOFT_RIGHT":                    2,
		"KEYCODE_HOME":                          3,
		"KEYCODE_BACK":                          4,
		"KEYCODE_CALL":                          5,
		"KEYCODE_ENDCALL":                       6,
		"KEYCODE_0":                             7,
		"KEYCODE_1":                             8,
		"KEYCODE_2":                             9,
		"KEYCODE_3":                             10,
		"KEYCODE_4":                             11,
		"KEYCODE_5":                             12,
		"KEYCODE_6":                             13,
		"KEYCODE_7":                             14,
		"KEYCODE_8":                             15,
		"KEYCODE_9":                             16,
		"KEYCODE_STAR":                          17,
		"KEYCODE_POUND":                         18,
		"KEYCODE_DPAD_UP":                       19,
		"KEYCODE_DPAD_DOWN":                     20,
		"KEYCODE_DPAD_LEFT":                     21,
		"KEYCODE_DPAD_RIGHT":                    22,
		"KEYCODE_DPAD_CENTER":                   23,
		"KEYCODE_VOLUME_UP":                     24,
		"KEYCODE_VOLUME_DOWN":                   25,
		"KEYCODE_POWER":                         26,
		"KEYCODE_CAMERA":                        27,
		"KEYCODE_CLEAR":                         28,
		"KEYCODE_A":                             29,
		"KEYCODE_B":                             30,
		"KEYCODE_C":                             31,
		"KEYCODE_D":                             32,
		"KEYCODE_E":                             33,
		"KEYCODE_F":                             34,
		"KEYCODE_G":                             35,
		"KEYCODE_H":                             36,
		"KEYCODE_I":                             37,
		"KEYCODE_J":                             38,
		"KEYCODE_K":                             39,
		"KEYCODE_L":                             40,
		"KEYCODE_M":                             41,
		"KEYCODE_N":                             42,
		"KEYCODE_O":                             43,
		"KEYCODE_P":                             44,
		"KEYCODE_Q":                             45,
		"KEYCODE_R":                             46,
		"KEYCODE_S":                             47,
		"KEYCODE_T":                             48,
		"KEYCODE_U":                             49,
		"KEYCODE_V":                             50,
		"KEYCODE_W":                             51,
		"KEYCODE_X":                             52,
		"KEYCODE_Y":                             53,
		"KEYCODE_Z":                             54,
		"KEYCODE_COMMA":                         55,
		"KEYCODE_PERIOD":                        56,
		"KEYCODE_ALT_LEFT":                      57,
		"KEYCODE_ALT_RIGHT":                     58,
		"KEYCODE_SHIFT_LEFT":                    59,
		"KEYCODE_SHIFT_RIGHT":                   60,
		"KEYCODE_TAB":                           61,
		"KEYCODE_SPACE":                         62,
		"KEYCODE_SYM":                           63,
		"KEYCODE_EXPLORER":                      64,
		"KEYCODE_ENVELOPE":                      65,
		"KEYCODE_ENTER":                         66,
		"KEYCODE_DEL":                           67,
		"KEYCODE_GRAVE":                         68,
		"KEYCODE_MINUS":                         69,
		"KEYCODE_EQUALS":                        70,
		"KEYCODE_LEFT_BRACKET":                  71,
		"KEYCODE_RIGHT_BRACKET":                 72,
		"KEYCODE_BACKSLASH":                     73,
		"KEYCODE_SEMICOLON":                     74,
		"KEYCODE_APOSTROPHE":                    75,
		"KEYCODE_SLASH":                         76,
		"KEYCODE_AT":                            77,
		"KEYCODE_NUM":                           78,
		"KEYCODE_HEADSETHOOK":                   79,
		"KEYCODE_FOCUS":                         80,
		"KEYCODE_PLUS":                          81,
		"KEYCODE_MENU":                          82,
		"KEYCODE_NOTIFICATION":                  83,
		"KEYCODE_SEARCH":                        84,
		"KEYCODE_MEDIA_PLAY_PAUSE":              85,
		"KEYCODE_MEDIA_STOP":                    86,
		"KEYCODE_MEDIA_NEXT":                    87,
		"KEYCODE_MEDIA_PREVIOUS":                88,
		"KEYCODE_MEDIA_REWIND":                  89,
		"KEYCODE_MEDIA_FAST_FORWARD":            90,
		"KEYCODE_MUTE":                          91,
		"KEYCODE_PAGE_UP":                       92,
		"KEYCODE_PAGE_DOWN":                     93,
		"KEYCODE_PICTSYMBOLS":                   94,
		"KEYCODE_SWITCH_CHARSET":                95,
		"KEYCODE_BUTTON_A":                      96,
		"KEYCODE_BUTTON_B":                      97,
		"KEYCODE_BUTTON_C":                      98,
		"KEYCODE_BUTTON_X":                      99,
		"KEYCODE_BUTTON_Y":                      100,
		"KEYCODE_BUTTON_Z":                      101,
		"KEYCODE_BUTTON_L1":                     102,
		"KEYCODE_BUTTON_R1":                     103,
		"KEYCODE_BUTTON_L2":                     104,
		"KEYCODE_BUTTON_R2":                     105,
		"KEYCODE_BUTTON_THUMBL":                 106,
		"KEYCODE_BUTTON_THUMBR":                 107,
		"KEYCODE_BUTTON_START":                  108,
		"KEYCODE_BUTTON_SELECT":                 109,
		"KEYCODE_BUTTON_MODE":                   110,
		"KEYCODE_ESCAPE":                        111,
		"KEYCODE_FORWARD_DEL":                   112,
		"KEYCODE_CTRL_LEFT":                     113,
		"KEYCODE_CTRL_RIGHT":                    114,
		"KEYCODE_CAPS_LOCK":                     115,
		"KEYCODE_SCROLL_LOCK":                   116,
		"KEYCODE_META_LEFT":                     117,
		"KEYCODE_META_RIGHT":                    118,
		"KEYCODE_FUNCTION":                      119,
		"KEYCODE_SYSRQ":                         120,
		"KEYCODE_BREAK":                         121,
		"KEYCODE_MOVE_HOME":                     122,
		"KEYCODE_MOVE_END":                      123,
		"KEYCODE_INSERT":                        124,
		"KEYCODE_FORWARD":                       125,
		"KEYCODE_MEDIA_PLAY":                    126,
		"KEYCODE_MEDIA_PAUSE":                   127,
		"KEYCODE_MEDIA_CLOSE":                   128,
		"KEYCODE_MEDIA_EJECT":                   129,
		"KEYCODE_MEDIA_RECORD":                  130,
		"KEYCODE_F1":                            131,
		"KEYCODE_F2":                            132,
		"KEYCODE_F3":                            133,
		"KEYCODE_F4":                            134,
		"KEYCODE_F5":                            135,
		"KEYCODE_F6":                            136,
		"KEYCODE_F7":                            137,
		"KEYCODE_F8":                            138,
		"KEYCODE_F9":                            139,
		"KEYCODE_F10":                           140,
		"KEYCODE_F11":                           141,
		"KEYCODE_F12":                           142,
		"KEYCODE_NUM_LOCK":                      143,
		"KEYCODE_NUMPAD_0":                      144,
		"KEYCODE_NUMPAD_1":                      145,
		"KEYCODE_NUMPAD_2":                      146,
		"KEYCODE_NUMPAD_3":                      147,
		"KEYCODE_NUMPAD_4":                      148,
		"KEYCODE_NUMPAD_5":                      149,
		"KEYCODE_NUMPAD_6":                      150,
		"KEYCODE_NUMPAD_7":                      151,
		"KEYCODE_NUMPAD_8":                      152,
		"KEYCODE_NUMPAD_9":                      153,
		"KEYCODE_NUMPAD_DIVIDE":                 154,
		"KEYCODE_NUMPAD_MULTIPLY":               155,
		"KEYCODE_NUMPAD_SUBTRACT":               156,
		"KEYCODE_NUMPAD_ADD":                    157,
		"KEYCODE_NUMPAD_DOT":                    158,
		"KEYCODE_NUMPAD_COMMA":                  159,
		"KEYCODE_NUMPAD_ENTER":                  160,
		"KEYCODE_NUMPAD_EQUALS":                 161,
		"KEYCODE_NUMPAD_LEFT_PAREN":             162,
		"KEYCODE_NUMPAD_RIGHT_PAREN":            163,
		"KEYCODE_VOLUME_MUTE":                   164,
		"KEYCODE_INFO":                          165,
		"KEYCODE_CHANNEL_UP":                    166,
		"KEYCODE_CHANNEL_DOWN":                  167,
		"KEYCODE_ZOOM_IN":                       168,
		"KEYCODE_ZOOM_OUT":                      169,
		"KEYCODE_TV":                            170,
		"KEYCODE_WINDOW":                        171,
		"KEYCODE_GUIDE":                         172,
		"KEYCODE_DVR":                           173,
		"KEYCODE_BOOKMARK":                      174,
		"KEYCODE_CAPTIONS":                      175,
		"KEYCODE_SETTINGS":                      176,
		"KEYCODE_TV_POWER":                      177,
		"KEYCODE_TV_INPUT":                      178,
		"KEYCODE_STB_POWER":                     179,
		"KEYCODE_STB_INPUT":                     180,
		"KEYCODE_AVR_POWER":                     181,
		"KEYCODE_AVR_INPUT":                     182,
		"KEYCODE_PROG_RED":                      183,
		"KEYCODE_PROG_GREEN":                    184,
		"KEYCODE_PROG_YELLOW":                   185,
		"KEYCODE_PROG_BLUE":                     186,
		"KEYCODE_APP_SWITCH":                    187,
		"KEYCODE_BUTTON_1":                      188,
		"KEYCODE_BUTTON_2":                      189,
		"KEYCODE_BUTTON_3":                      190,
		"KEYCODE_BUTTON_4":                      191,
		"KEYCODE_BUTTON_5":                      192,
		"KEYCODE_BUTTON_6":                      193,
		"KEYCODE_BUTTON_7":                      194,
		"KEYCODE_BUTTON_8":                      195,
		"KEYCODE_BUTTON_9":                      196,
		"KEYCODE_BUTTON_10":                     197,
		"KEYCODE_BUTTON_11":                     198,
		"KEYCODE_BUTTON_12":                     199,
		"KEYCODE_BUTTON_13":                     200,
		"KEYCODE_BUTTON_14":                     201,
		"KEYCODE_BUTTON_15":                     202,
		"KEYCODE_BUTTON_16":                     203,
		"KEYCODE_LANGUAGE_SWITCH":               204,
		"KEYCODE_MANNER_MODE":                   205,
		"KEYCODE_3D_MODE":                       206,
		"KEYCODE_CONTACTS":                      207,
		"KEYCODE_CALENDAR":                      208,
		"KEYCODE_MUSIC":                         209,
		"KEYCODE_CALCULATOR":                    210,
		"KEYCODE_ZENKAKU_HANKAKU":               211,
		"KEYCODE_EISU":                          212,
		"KEYCODE_MUHENKAN":                      213,
		"KEYCODE_HENKAN":                        214,
		"KEYCODE_KATAKANA_HIRAGANA":             215,
		"KEYCODE_YEN":                           216,
		"KEYCODE_RO":                            217,
		"KEYCODE_KANA":                          218,
		"KEYCODE_ASSIST":                        219,
		"KEYCODE_BRIGHTNESS_DOWN":               220,
		"KEYCODE_BRIGHTNESS_UP":                 221,
		"KEYCODE_MEDIA_AUDIO_TRACK":             222,
		"KEYCODE_SLEEP":                         223,
		"KEYCODE_WAKEUP":                        224,
		"KEYCODE_PAIRING":                       225,
		"KEYCODE_MEDIA_TOP_MENU":                226,
		"KEYCODE_11":                            227,
		"KEYCODE_12":                            228,
		"KEYCODE_LAST_CHANNEL":                  229,
		"KEYCODE_TV_DATA_SERVICE":               230,
		"KEYCODE_VOICE_ASSIST":                  231,
		"KEYCODE_TV_RADIO_SERVICE":              232,
		"KEYCODE_TV_TELETEXT":                   233,
		"KEYCODE_TV_NUMBER_ENTRY":               234,
		"KEYCODE_TV_TERRESTRIAL_ANALOG":         235,
		"KEYCODE_TV_TERRESTRIAL_DIGITAL":        236,
		"KEYCODE_TV_SATELLITE":                  237,
		"KEYCODE_TV_SATELLITE_BS":               238,
		"KEYCODE_TV_SATELLITE_CS":               239,
		"KEYCODE_TV_SATELLITE_SERVICE":          240,
		"KEYCODE_TV_NETWORK":                    241,
		"KEYCODE_TV_ANTENNA_CABLE":              242,
		"KEYCODE_TV_INPUT_HDMI_1":               243,
		"KEYCODE_TV_INPUT_HDMI_2":               244,
		"KEYCODE_TV_INPUT_HDMI_3":               245,
		"KEYCODE_TV_INPUT_HDMI_4":               246,
		"KEYCODE_TV_INPUT_COMPOSITE_1":          247,
		"KEYCODE_TV_INPUT_COMPOSITE_2":          248,
		"KEYCODE_TV_INPUT_COMPONENT_1":          249,
		"KEYCODE_TV_INPUT_COMPONENT_2":          250,
		"KEYCODE_TV_INPUT_VGA_1":                251,
		"KEYCODE_TV_AUDIO_DESCRIPTION":          252,
		"KEYCODE_TV_AUDIO_DESCRIPTION_MIX_UP":   253,
		"KEYCODE_TV_AUDIO_DESCRIPTION_MIX_DOWN": 254,
		"KEYCODE_TV_ZOOM_MODE":                  255,
		"KEYCODE_TV_CONTENTS_MENU":              256,
		"KEYCODE_TV_MEDIA_CONTEXT_MENU":         257,
		"KEYCODE_TV_TIMER_PROGRAMMING":          258,
		"KEYCODE_HELP":                          259,
		"KEYCODE_NAVIGATE_PREVIOUS":             260,
		"KEYCODE_NAVIGATE_NEXT":                 261,
		"KEYCODE_NAVIGATE_IN":                   262,
		"KEYCODE_NAVIGATE_OUT":                  263,
		"KEYCODE_STEM_PRIMARY":                  264,
		"KEYCODE_STEM_1":                        265,
		"KEYCODE_STEM_2":                        266,
		"KEYCODE_STEM_3":                        267,
		"KEYCODE_DPAD_UP_LEFT":                  268,
		"KEYCODE_DPAD_DOWN_LEFT":                269,
		"KEYCODE_DPAD_UP_RIGHT":                 270,
		"KEYCODE_DPAD_DOWN_RIGHT":               271,
		"KEYCODE_MEDIA_SKIP_FORWARD":            272,
		"KEYCODE_MEDIA_SKIP_BACKWARD":           273,
		"KEYCODE_MEDIA_STEP_FORWARD":            274,
		"KEYCODE_MEDIA_STEP_BACKWARD":           275,
		"KEYCODE_SOFT_SLEEP":                    276,
		"KEYCODE_CUT":                           277,
		"KEYCODE_COPY":                          278,
		"KEYCODE_PASTE":                         279,
		"KEYCODE_SYSTEM_NAVIGATION_UP":          280,
		"KEYCODE_SYSTEM_NAVIGATION_DOWN":        281,
		"KEYCODE_SYSTEM_NAVIGATION_LEFT":        282,
		"KEYCODE_SYSTEM_NAVIGATION_RIGHT":       283,
		"KEYCODE_ALL_APPS":                      284,
		"KEYCODE_REFRESH":                       285,
		"KEYCODE_THUMBS_UP":                     286,
		"KEYCODE_THUMBS_DOWN":                   287,
		"KEYCODE_PROFILE_SWITCH":                288,
		"KEYCODE_VIDEO_APP_1":                   289,
		"KEYCODE_VIDEO_APP_2":                   290,
		"KEYCODE_VIDEO_APP_3":                   291,
		"KEYCODE_VIDEO_APP_4":                   292,
		"KEYCODE_VIDEO_APP_5":                   293,
		"KEYCODE_VIDEO_APP_6":                   294,
		"KEYCODE_VIDEO_APP_7":                   295,
		"KEYCODE_VIDEO_APP_8":                   296,
		"KEYCODE_FEATURED_APP_1":                297,
		"KEYCODE_FEATURED_APP_2":                298,
		"KEYCODE_FEATURED_APP_3":                299,
		"KEYCODE_FEATURED_APP_4":                300,
		"KEYCODE_DEMO_APP_1":                    301,
		"KEYCODE_DEMO_APP_2":                    302,
		"KEYCODE_DEMO_APP_3":                    303,
		"KEYCODE_DEMO_APP_4":                    304,
	}
)
