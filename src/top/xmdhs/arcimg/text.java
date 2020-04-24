package top.xmdhs.arcimg;


public class text {
    public static String song() {
        String asong;
        String difficulty = a.json.value.get(0).value.friends.get(0).recent_score.get(0).difficulty;
        if (difficulty.equals("0")) {
            difficulty = "PST";
        }
        if (difficulty.equals("1")) {
            difficulty = "PRS";
        }
        if (difficulty.equals("2")) {
            difficulty = "FTR";
        }
        asong = a.json.value.get(0).value.friends.get(0).recent_score.get(0).song_id + " (" + difficulty + ")";
        return asong;
    }

    public static String type() {
        String type = a.json.value.get(0).value.friends.get(0).recent_score.get(0).clear_type;
        if (type.equals("0")) {
            return "Track Lost";
        }
        if (type.equals("1")) {
            return "Normal Clear";
        }
        if (type.equals("2")) {
            return "Full Recall";
        }
        if (type.equals("3")) {
            return "Pure Memory";
        }
        if (type.equals("4")) {
            return "Easy Clear";
        }
        if (type.equals("5")) {
            return "Hard Clear";
        }
        return "";
    }

    public static String time() {
        return convertTimeToFormat(a.json.value.get(0).value.friends.get(0).recent_score.get(0).time_played / 1000);
    }

    public static String pure() {
        return "PURE: " + a.json.value.get(0).value.friends.get(0).recent_score.get(0).perfect_count +
                "(" + a.json.value.get(0).value.friends.get(0).recent_score.get(0).shiny_perfect_count + ")";
    }

    public static String far() {
        return "FAR: " + a.json.value.get(0).value.friends.get(0).recent_score.get(0).near_count;
    }

    public static String lost() {
        return "LOST: " + a.json.value.get(0).value.friends.get(0).recent_score.get(0).miss_count;
    }

    public static String ptt() {
        int l = a.json.value.get(0).value.friends.get(0).rating.length();
        if (l == 4) {
            String ptt = a.json.value.get(0).value.friends.get(0).rating.substring(0, 2) + "." +
                    a.json.value.get(0).value.friends.get(0).rating.substring(2);
            return "PTT: " + ptt;
        }
        if (l == 3) {
            String ptt = a.json.value.get(0).value.friends.get(0).rating.substring(0, 1) + "." +
                    a.json.value.get(0).value.friends.get(0).rating.substring(1);
            return "PTT: " + ptt;
        }
        return "PTT: " + a.json.value.get(0).value.friends.get(0).rating;
    }

    public static String rating() {
        if (a.json.value.get(0).value.friends.get(0).recent_score.get(0).rating.length() > 7) {
            return "Result rating: " + a.json.value.get(0).value.friends.get(0).recent_score.get(0).rating.substring(0, 7);
        } else {
            return "Result rating: " + a.json.value.get(0).value.friends.get(0).recent_score.get(0).rating;
        }
    }

    public static String convertTimeToFormat(long timeStamp) {
        long curTime = System.currentTimeMillis() / (long) 1000;
        long time = curTime - timeStamp;

        if (time < 60 && time >= 0) {
            return "now";
        } else if (time >= 60 && time < 3600) {
            if (time / 60 == 1) {
                return time / 60 + " minute ago";
            } else {
                return time / 60 + " minutes ago";
            }
        } else if (time >= 3600 && time < 3600 * 24) {
            if (time / 3600 == 1) {
                return time / 3600 + " hour ago";
            } else {
                return time / 3600 + " hours ago";
            }
        } else if (time >= 3600 * 24 && time < 3600 * 24 * 30) {
            if (time / 3600 / 24 == 1) {
                return time / 3600 / 24 + " day ago";
            } else {
                return time / 3600 / 24 + " days ago";
            }
        } else if (time >= 3600 * 24 * 30 && time < 3600 * 24 * 30 * 12) {
            if (time / 3600 / 24 / 30 == 1) {
                return time / 3600 / 24 / 30 + " month ago";
            } else {
                return time / 3600 / 24 / 30 + " months ago";
            }
        } else if (time >= 3600 * 24 * 30 * 12) {
            if (time / 3600 / 24 / 30 / 12 == 1) {
                return time / 3600 / 24 / 30 / 12 + " year ago";
            } else {
                return time / 3600 / 24 / 30 / 12 + " years ago";
            }
        } else {
            return "now";
        }
    }
}

