//go:build darwin

#ifndef _events_h
#define _events_h

extern void processApplicationEvent(unsigned int, void* data);
extern void processWindowEvent(unsigned int, unsigned int);

#define EventApplicationDidBecomeActive 1024
#define EventApplicationDidChangeBackingProperties 1025
#define EventApplicationDidChangeEffectiveAppearance 1026
#define EventApplicationDidChangeIcon 1027
#define EventApplicationDidChangeOcclusionState 1028
#define EventApplicationDidChangeScreenParameters 1029
#define EventApplicationDidChangeStatusBarFrame 1030
#define EventApplicationDidChangeStatusBarOrientation 1031
#define EventApplicationDidFinishLaunching 1032
#define EventApplicationDidHide 1033
#define EventApplicationDidResignActiveNotification 1034
#define EventApplicationDidUnhide 1035
#define EventApplicationDidUpdate 1036
#define EventApplicationWillBecomeActive 1037
#define EventApplicationWillFinishLaunching 1038
#define EventApplicationWillHide 1039
#define EventApplicationWillResignActive 1040
#define EventApplicationWillTerminate 1041
#define EventApplicationWillUnhide 1042
#define EventApplicationWillUpdate 1043
#define EventApplicationDidChangeTheme 1044
#define EventApplicationShouldHandleReopen 1045
#define EventWindowDidBecomeKey 1046
#define EventWindowDidBecomeMain 1047
#define EventWindowDidBeginSheet 1048
#define EventWindowDidChangeAlpha 1049
#define EventWindowDidChangeBackingLocation 1050
#define EventWindowDidChangeBackingProperties 1051
#define EventWindowDidChangeCollectionBehavior 1052
#define EventWindowDidChangeEffectiveAppearance 1053
#define EventWindowDidChangeOcclusionState 1054
#define EventWindowDidChangeOrderingMode 1055
#define EventWindowDidChangeScreen 1056
#define EventWindowDidChangeScreenParameters 1057
#define EventWindowDidChangeScreenProfile 1058
#define EventWindowDidChangeScreenSpace 1059
#define EventWindowDidChangeScreenSpaceProperties 1060
#define EventWindowDidChangeSharingType 1061
#define EventWindowDidChangeSpace 1062
#define EventWindowDidChangeSpaceOrderingMode 1063
#define EventWindowDidChangeTitle 1064
#define EventWindowDidChangeToolbar 1065
#define EventWindowDidChangeVisibility 1066
#define EventWindowDidDeminiaturize 1067
#define EventWindowDidEndSheet 1068
#define EventWindowDidEnterFullScreen 1069
#define EventWindowDidEnterVersionBrowser 1070
#define EventWindowDidExitFullScreen 1071
#define EventWindowDidExitVersionBrowser 1072
#define EventWindowDidExpose 1073
#define EventWindowDidFocus 1074
#define EventWindowDidMiniaturize 1075
#define EventWindowDidMove 1076
#define EventWindowDidOrderOffScreen 1077
#define EventWindowDidOrderOnScreen 1078
#define EventWindowDidResignKey 1079
#define EventWindowDidResignMain 1080
#define EventWindowDidResize 1081
#define EventWindowDidUpdate 1082
#define EventWindowDidUpdateAlpha 1083
#define EventWindowDidUpdateCollectionBehavior 1084
#define EventWindowDidUpdateCollectionProperties 1085
#define EventWindowDidUpdateShadow 1086
#define EventWindowDidUpdateTitle 1087
#define EventWindowDidUpdateToolbar 1088
#define EventWindowDidUpdateVisibility 1089
#define EventWindowShouldClose 1090
#define EventWindowWillBecomeKey 1091
#define EventWindowWillBecomeMain 1092
#define EventWindowWillBeginSheet 1093
#define EventWindowWillChangeOrderingMode 1094
#define EventWindowWillClose 1095
#define EventWindowWillDeminiaturize 1096
#define EventWindowWillEnterFullScreen 1097
#define EventWindowWillEnterVersionBrowser 1098
#define EventWindowWillExitFullScreen 1099
#define EventWindowWillExitVersionBrowser 1100
#define EventWindowWillFocus 1101
#define EventWindowWillMiniaturize 1102
#define EventWindowWillMove 1103
#define EventWindowWillOrderOffScreen 1104
#define EventWindowWillOrderOnScreen 1105
#define EventWindowWillResignMain 1106
#define EventWindowWillResize 1107
#define EventWindowWillUnfocus 1108
#define EventWindowWillUpdate 1109
#define EventWindowWillUpdateAlpha 1110
#define EventWindowWillUpdateCollectionBehavior 1111
#define EventWindowWillUpdateCollectionProperties 1112
#define EventWindowWillUpdateShadow 1113
#define EventWindowWillUpdateTitle 1114
#define EventWindowWillUpdateToolbar 1115
#define EventWindowWillUpdateVisibility 1116
#define EventWindowWillUseStandardFrame 1117
#define EventMenuWillOpen 1118
#define EventMenuDidOpen 1119
#define EventMenuDidClose 1120
#define EventMenuWillSendAction 1121
#define EventMenuDidSendAction 1122
#define EventMenuWillHighlightItem 1123
#define EventMenuDidHighlightItem 1124
#define EventMenuWillDisplayItem 1125
#define EventMenuDidDisplayItem 1126
#define EventMenuWillAddItem 1127
#define EventMenuDidAddItem 1128
#define EventMenuWillRemoveItem 1129
#define EventMenuDidRemoveItem 1130
#define EventMenuWillBeginTracking 1131
#define EventMenuDidBeginTracking 1132
#define EventMenuWillEndTracking 1133
#define EventMenuDidEndTracking 1134
#define EventMenuWillUpdate 1135
#define EventMenuDidUpdate 1136
#define EventMenuWillPopUp 1137
#define EventMenuDidPopUp 1138
#define EventMenuWillSendActionToItem 1139
#define EventMenuDidSendActionToItem 1140
#define EventWebViewDidStartProvisionalNavigation 1141
#define EventWebViewDidReceiveServerRedirectForProvisionalNavigation 1142
#define EventWebViewDidFinishNavigation 1143
#define EventWebViewDidCommitNavigation 1144
#define EventWindowFileDraggingEntered 1145
#define EventWindowFileDraggingPerformed 1146
#define EventWindowFileDraggingExited 1147

#define MAX_EVENTS 1148


#endif