package adventofcode.year2018;

import adventofcode.BaseTest;
import lombok.Data;
import org.apache.commons.lang3.tuple.Pair;
import org.junit.Test;

import java.security.Guard;
import java.time.Duration;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.Period;
import java.time.format.DateTimeFormatter;
import java.util.*;

public class Day4 extends BaseTest {

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2018/day4/input.txt");
		List<Entry> entries = new ArrayList<>();
		for (final String input : data) {
			entries.add(new Entry(input));
		}
		entries.sort(Comparator.comparing(o -> o.getDate()));

		Iterator<Entry> it =  entries.iterator();
		List<Guard> guards = new ArrayList<>();
		Guard currGuard  = null;
		while (it.hasNext()){
			Entry curr = it.next();
			if (curr.isGuard()){
				currGuard = new Guard(curr);
				guards.add(currGuard);
			}else{
				currGuard.addShift(curr, it.next());
			}
		}
		Map<String, List<Guard>> guardsMap= new HashMap<>();
		for (Guard guard : guards){
			if (guardsMap.containsKey(guard.id)){
				guardsMap.get(guard.id).add(guard);
			}else{
				List<Guard> curL = new ArrayList<>();
				curL.add(guard);
				guardsMap.put(guard.id, curL);
			}
		}
		Map<String, Long> sleepTime = new HashMap<>();

		for (Map.Entry<String, List<Guard>> entry : guardsMap.entrySet()) {
			long currSleep = 0l;
			for (Guard guard : entry.getValue()){
				currSleep +=guard.getSleepTime();
			}
			sleepTime.put(entry.getKey(), currSleep);
		}
		String maxSleepId = "";
		long maxSleepTime = 0l;

		for (Map.Entry<String,Long> entry : sleepTime.entrySet()) {
			if (maxSleepTime<entry.getValue()){
				maxSleepId = entry.getKey();
				maxSleepTime = entry.getValue();
			}
		}
		System.out.println(maxSleepId + " "  + maxSleepTime);
		int mostSleepMin = 0;
		int mostSleepMinCount = 0;
		List<Guard> mostSleepGuard = guardsMap.get(maxSleepId);
		for (int i=0; i<60; i++){
			int sleepCount = 0;
			for (Guard guard: mostSleepGuard){
				if (guard.sleepAtMinute(i)){
					sleepCount ++;
				}
			}
			if (sleepCount>mostSleepMinCount){
				mostSleepMin = i;
				mostSleepMinCount = sleepCount;
				System.out.println("mostSleepMin: " + mostSleepMin + "  mostSleepMinCount" + mostSleepMinCount);
			}
		}
		int guardId = Integer.parseInt(maxSleepId.substring(maxSleepId.indexOf("#")+1, maxSleepId.indexOf(" begins")));
		System.out.println(mostSleepMin);
		System.out.println(mostSleepMin*guardId);

	}


	@Data
	class Guard {
		String id;
		List<Shift> shifts = new ArrayList<>();
		public Guard(Entry name){
			this.id = name.action;
		}
		public void addShift(Entry start, Entry end){
			this.shifts.add(new Shift(start,end));
		}
		public long getSleepTime(){
			long sleep  = 0l;
			for (Shift shift:shifts){
				sleep += shift.minInBetween();
			}
			return sleep;
		}

		public boolean sleepAtMinute(int min){
			for (Shift shift : shifts){
				if(shift.sleepAtMinute(min)){
					return true;
				}
			}
				return false;
		}
	}

	@Data
	class Shift{
		LocalDateTime falls;
		LocalDateTime wakes;
		public Shift(Entry start, Entry end){
			this.falls  = start.date;
			this.wakes  = end.date;
		}

		public long minInBetween(){
			return  Duration.between(falls, wakes).toMinutes();
		}
		public boolean sleepAtMinute(int min){
			return  falls.getMinute() <= min && min < wakes.getMinute();
		}

	}

	@Data
	class Entry{
		LocalDateTime date;
		String action;
		public Entry(String input){
			String time = input.substring(1, input.indexOf(']'));
			DateTimeFormatter formatter = DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm");
			this.date  = LocalDateTime.parse(time, formatter);
			this.action = input.substring(input.indexOf(']') + 2);
		}
		public boolean isGuard(){
			return action.startsWith("Guard");
		}
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2018/day4/input.txt");
		List<Entry> entries = new ArrayList<>();
		for (final String input : data) {
			entries.add(new Entry(input));
		}
		entries.sort(Comparator.comparing(o -> o.getDate()));

		Iterator<Entry> it =  entries.iterator();
		List<Guard> guards = new ArrayList<>();
		Guard currGuard  = null;
		while (it.hasNext()){
			Entry curr = it.next();
			if (curr.isGuard()){
				currGuard = new Guard(curr);
				guards.add(currGuard);
			}else{
				currGuard.addShift(curr, it.next());
			}
		}
		Map<String, List<Guard>> guardsMap= new HashMap<>();
		for (Guard guard : guards){
			if (guardsMap.containsKey(guard.id)){
				guardsMap.get(guard.id).add(guard);
			}else{
				List<Guard> curL = new ArrayList<>();
				curL.add(guard);
				guardsMap.put(guard.id, curL);
			}
		}
		Map<String, Long> sleepTime = new HashMap<>();

		for (Map.Entry<String, List<Guard>> entry : guardsMap.entrySet()) {
			long currSleep = 0l;
			for (Guard guard : entry.getValue()){
				currSleep +=guard.getSleepTime();
			}
			sleepTime.put(entry.getKey(), currSleep);
		}

		int mostSleepMin = 0;
		int mostSleepMinCount = 0;
		String maxSleepGardsId = "";


		for (Map.Entry<String, List<Guard>> entry : guardsMap.entrySet()) {
			String curSleepGardsId  = entry.getKey();
			List<Guard> mostSleepGuard = entry.getValue();
			for (int i=0; i<60; i++){
				int sleepCount = 0;
				for (Guard guard: mostSleepGuard){
					if (guard.sleepAtMinute(i)){
						sleepCount ++;
					}
				}
				if (sleepCount>mostSleepMinCount){
					mostSleepMin = i;
					mostSleepMinCount = sleepCount;
					maxSleepGardsId = curSleepGardsId;
					System.out.println( " maxSleepGardsId" + maxSleepGardsId + " mostSleepMin: " + mostSleepMin + "  mostSleepMinCount " + mostSleepMinCount);
				}
			}
		}
		int guardId = Integer.parseInt(maxSleepGardsId.substring(maxSleepGardsId.indexOf("#")+1, maxSleepGardsId.indexOf(" begins")));
		System.out.println(mostSleepMin);
		System.out.println(mostSleepMin*guardId);

	}

}
