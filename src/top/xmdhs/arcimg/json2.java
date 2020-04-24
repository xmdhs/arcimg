package top.xmdhs.arcimg;

import com.google.gson.Gson;

import java.util.List;

public class json2 {
    public static void main(String[] args) {
        String jj = "{\"success\":true,\"value\":[{\"id\":0,\"value\":{\"is_aprilfools\":false,\"curr_available_maps\":[],\"character_stats\":[{\"is_uncapped\":false,\"uncap_cores\":[{\"_id\":\"5e7ac3994c53bf01b7d1614f\",\"core_type\":\"core_hollow\",\"amount\":25},{\"_id\":\"5e7ac3994c53bf01b7d1614e\",\"core_type\":\"core_desolate\",\"amount\":5}],\"char_type\":1,\"skill_id_uncap\":\"\",\"skill_requires_uncap\":false,\"skill_unlock_level\":0,\"skill_id\":\"gauge_easy\",\"overdrive\":35.01516256013996,\"prog\":35.01516256013996,\"frag\":55.01341303396997,\"level_exp\":50,\"exp\":93.83728,\"level\":2,\"name\":\"hikari\",\"character_id\":0},{\"is_uncapped\":false,\"uncap_cores\":[{\"_id\":\"5e7ac3994c53bf01b7d16151\",\"core_type\":\"core_desolate\",\"amount\":25},{\"_id\":\"5e7ac3994c53bf01b7d16150\",\"core_type\":\"core_hollow\",\"amount\":5}],\"char_type\":0,\"skill_id_uncap\":\"\",\"skill_requires_uncap\":false,\"skill_unlock_level\":0,\"skill_id\":\"\",\"overdrive\":62.46464499198134,\"prog\":62.46464499198134,\"frag\":62.46464499198134,\"level_exp\":900,\"exp\":927.1788,\"level\":9,\"name\":\"tairitsu\",\"character_id\":1},{\"is_previewable\":true,\"uncap_cores\":[{\"_id\":\"5e7ac3994c53bf01b7d16153\",\"core_type\":\"core_crimson\",\"amount\":25},{\"_id\":\"5e7ac3994c53bf01b7d16152\",\"core_type\":\"core_hollow\",\"amount\":5}],\"char_type\":0,\"skill_id_uncap\":\"\",\"skill_requires_uncap\":true,\"skill_unlock_level\":20,\"skill_id\":\"frags_kou\",\"overdrive\":70,\"prog\":70,\"frag\":90,\"level_exp\":10000,\"exp\":10000,\"level\":20,\"name\":\"kou\",\"character_id\":2},{\"is_previewable\":true,\"uncap_cores\":[],\"char_type\":0,\"skill_id_uncap\":\"\",\"skill_requires_uncap\":false,\"skill_unlock_level\":0,\"skill_id\":\"\",\"overdrive\":75,\"prog\":75,\"frag\":75,\"level_exp\":10000,\"exp\":10000,\"level\":20,\"name\":\"sapphire\",\"character_id\":3},{\"is_previewable\":true,\"uncap_cores\":[{\"_id\":\"5e7ac3994c53bf01b7d16155\",\"core_type\":\"core_ambivalent\",\"amount\":25},{\"_id\":\"5e7ac3994c53bf01b7d16154\",\"core_type\":\"core_desolate\",\"amount\":5}],\"char_type\":0,\"skill_id_uncap\":\"visual_ink\",\"skill_requires_uncap\":false,\"skill_unlock_level\":8,\"skill_id\":\"note_mirror\",\"overdrive\":70,\"prog\":90,\"frag\":70,\"level_exp\":10000,\"exp\":10000,\"level\":20,\"name\":\"lethe\",\"character_id\":4}],\"friends\":[{\"is_mutual\":false,\"is_char_uncapped_override\":false,\"is_char_uncapped\":true,\"is_skill_sealed\":false,\"rating\":961,\"join_date\":1572538347498,\"character\":2,\"recent_score\":[{\"rating\":0,\"modifier\":2,\"time_played\":1587250285320,\"health\":-1,\"best_clear_type\":0,\"clear_type\":0,\"miss_count\":476,\"near_count\":0,\"perfect_count\":0,\"shiny_perfect_count\":0,\"score\":0,\"difficulty\":0,\"song_id\":\"pragmatism\"}],\"name\":\"xmdhs\",\"user_id\":2476615}],\"settings\":{\"is_hide_rating\":false,\"favorite_character\":-1,\"max_stamina_notification_enabled\":false},\"user_id\":2550615,\"name\":\"yneit\",\"user_code\":\"073158500\",\"display_name\":\"yneit\",\"ticket\":0,\"character\":1,\"is_locked_name_duplicate\":false,\"is_skill_sealed\":false,\"current_map\":\"tairitsu_tech\",\"prog_boost\":0,\"next_fragstam_ts\":-1,\"max_stamina_ts\":1581695872393,\"stamina\":9,\"world_songs\":[\"babaroque\",\"shadesoflight\"],\"singles\":[],\"packs\":[],\"characters\":[0,1],\"cores\":[{\"core_type\":\"core_generic\",\"amount\":1,\"_id\":\"5e465757b2a9cd175fa14914\"}],\"recent_score\":[{\"song_id\":\"chronostasis\",\"difficulty\":2,\"score\":9443889,\"shiny_perfect_count\":658,\"perfect_count\":831,\"near_count\":68,\"miss_count\":17,\"clear_type\":1,\"best_clear_type\":1,\"health\":100,\"time_played\":1581692410858,\"modifier\":0}],\"max_friend\":10,\"rating\":837,\"join_date\":1574176396254}}]}\n";
        Gson gson = new Gson();
        Json j = gson.fromJson(jj, Json.class);
        System.out.println(j.value.get(0).value.friends.get(0).name);
        System.out.println(j.value.get(0).value.friends.get(0).rating);
        System.out.println(j.value.get(0).value.friends.get(0).recent_score.get(0).song_id);
    }
    public static Json json2class(String json){
        Gson gson = new Gson();
        return gson.fromJson(json, Json.class);
    }
}

class Json {
    public List<Value> value;
    public static class Value {
        public VAlue value;
        public static class VAlue {
            public List<Friends> friends;
            public static class Friends{
                public String name;//名字
                public String rating;//ptt
                public List<Recent> recent_score;
                public static class Recent{
                    public String clear_type; //通过类型？ 0 大概是没通过, 1 是普通，5 是困难 ,3 是 p ，简单是 4，full recall 是 2.
                    public String miss_count; //lost
                    public String near_count; //far
                    public String perfect_count; //pure
                    public String shiny_perfect_count; //大 p
                    public String score; //分数
                    public String rating; //评价
                    public String difficulty; //难度，0 为过去，1为现在，2 为未来。
                    public String song_id; //曲目
                    public Long time_played; //上次游玩时间，为 unix 时间戳。（毫秒
                }
            }
        }
    }

}
