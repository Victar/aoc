package adventofcode.year2020;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;

public class Day20 extends BaseTest {

	@Test public void singleCheck() {

	}

	@Test
	public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day20/input.txt");
		final List<Tile> tiles = new ArrayList<>();
		Tile currentTile = new Tile();
		for (String input : data) {
			if (StringUtils.isEmpty(input)) {
				tiles.add(currentTile);
				currentTile = new Tile();
			}else{
				Pattern p = Pattern.compile("\\d+");
		        Matcher m = p.matcher(input);
		        if(m.find()) {
		            currentTile.setId(Integer.parseInt(m.group()));
		        }else{
		        	currentTile.addDataLine(input);
		        }

			}
		}
		tiles.add(currentTile);
		for (Tile tile: tiles) {
			tile.initTiles();
			tile.printTile();
		}
		Map<Integer, Integer> tileBorderCount = new HashMap<>();
		for (Tile tile: tiles) {
			for (int i: tile.borderAll){
				Integer mapI =  tileBorderCount.get(i);
				if (mapI != null) {
					tileBorderCount.put(i, mapI+1);
				}else{
					tileBorderCount.put(i, 1);
				}
			}
		}
		List<Integer> singleBorder = new ArrayList<>();
		for (Map.Entry<Integer, Integer> entry : tileBorderCount.entrySet()) {
			if (entry.getValue() ==1){
				singleBorder.add(entry.getKey());
			}
		}
		long result = 1;
		for (Tile tile: tiles) {
			int sbCount = tile.singleBorderCount(singleBorder);
			if (sbCount >2) {
				System.out.println("sbCount" + sbCount);
				result*=tile.id;
				tile.printTile();
			}
		}
		System.out.println(tileBorderCount);
		System.out.println("Result: " + result);

	}

	@Test
	public void runGold() throws Exception{
		final ArrayList<String> data = readStringFromFile("year2020/day20/input.txt");
		final List<Tile> tiles = new ArrayList<>();
		Tile currentTile = new Tile();
		for (String input : data) {
			if (StringUtils.isEmpty(input)) {
				tiles.add(currentTile);
				currentTile = new Tile();
			}else{
				Pattern p = Pattern.compile("\\d+");
		        Matcher m = p.matcher(input);
		        if(m.find()) {
		            currentTile.setId(Integer.parseInt(m.group()));
		        }else{
		        	currentTile.addDataLine(input);
		        }

			}
		}
		tiles.add(currentTile);
		for (Tile tile: tiles) {
			tile.initTiles();
		}
		Map<Integer, Integer> tileBorderCount = new HashMap<>();
		for (Tile tile: tiles) {
			for (int i: tile.borderAll){
				Integer mapI =  tileBorderCount.get(i);
				if (mapI != null) {
					tileBorderCount.put(i, mapI+1);
				}else{
					tileBorderCount.put(i, 1);
				}
			}
		}
		List<Integer> singleBorder = new ArrayList<>();
		for (Map.Entry<Integer, Integer> entry : tileBorderCount.entrySet()) {
			if (entry.getValue() ==1){
				singleBorder.add(entry.getKey());
			}
		}
		int totalCount = 0;
		for (Tile tile: tiles) {
			totalCount+=tile.squareBorderCount(singleBorder);

		}
		int MONSTER_GUESS = 40;

		for (int i = 0; i < MONSTER_GUESS && i<10; i++) {
			System.out.println(totalCount - 15*(MONSTER_GUESS +i));
//			System.out.println(totalCount - 15*(MONSTER_GUESS-i));

		}
	}



	@Data
	class Tile{
		int id;
		int[] border = null;
		int[] borderFlipped = null;
		int[] borderAll = null;

		ArrayList<String> data = new ArrayList<String>();
		public void addDataLine(final String dataLine){
			data.add(dataLine);
		}
		public void initTiles() {
			initBorders();
		}


		public int squareBorderCount(List<Integer> singleBorder){
			int count = 0;
			for (int i=1; i<data.size()-1; i++){
				String substring = data.get(i).substring(1,data.get(i).length()-1);
//				System.out.println(substring);
				count+=StringUtils.countMatches(substring,"#");
			}
			int singleBorderCount = singleBorderCount(singleBorder);
			for (int i = 0; i < singleBorder.size(); i++) {
				if (contains(this.border, singleBorder.get(i))) {
					int borderCount = singleBorder.get(i);
					if (singleBorderCount>2){
//						System.out.println(borderCount + " ->" +  Integer.toBinaryString(borderCount ));
					}
				}
			}
			return count;
		}


		public int singleBorderCount(List<Integer> border){
			int count = 0;
			for (Integer b : border){
				if (containsBorder(b)){
					count++;
				}
			}
			return count;
		}

		public boolean containsBorder(final int border){
			return contains(this.borderAll, border);
		}
		public void printTile(){
			System.out.println("Tile " + id + ":");
			for (String s: data){
				System.out.println(s);
			}
			System.out.println(Arrays.toString(border));
			System.out.println(Arrays.toString(borderFlipped));
			System.out.println(Arrays.toString(borderAll));
			System.out.println("");
		}

		private  boolean contains(final int[] arr, final int key) {
		    return Arrays.stream(arr).anyMatch(i -> i == key);
		}

		public void initBorders(){
			if (border == null){
				StringBuilder sbX2 = new StringBuilder();
				StringBuilder sbX4 = new StringBuilder();
				for (int i=0; i< data.size(); i++){
					sbX4.append(data.get(i).charAt(0));
					sbX2.append(data.get(i).charAt(data.size()-1));
				}
				String s1 = StringUtils.replaceEach(data.get(0), new String[] { ".", "#" }, new String[] { "0", "1" });
				String s3 = StringUtils.replaceEach(StringUtils.reverse(data.get(data.size()-1)), new String[] { ".", "#" }, new String[] { "0", "1" });
				String s2 = StringUtils.replaceEach(sbX2.toString(), new String[] { ".", "#" }, new String[] { "0", "1" });
				String s4 = StringUtils.replaceEach(sbX4.reverse().toString(), new String[] { ".", "#" }, new String[] { "0", "1" });


				int x1= Integer.parseInt(s1, 2);
				int x3= Integer.parseInt(s3, 2);
				int x2= Integer.parseInt(s2, 2);
				int x4= Integer.parseInt(s4, 2);
				border = new int []{x1,x2,x3,x4};
				int xf1= Integer.parseInt(StringUtils.reverse(s1), 2);
				int xf3= Integer.parseInt(StringUtils.reverse(s3), 2);
				int xf2= Integer.parseInt(StringUtils.reverse(s2), 2);
				int xf4= Integer.parseInt(StringUtils.reverse(s4), 2);

				borderFlipped = new int []{xf1,xf2,xf3,xf4};
				borderAll = new int []{x1,x2,x3,x4,xf1,xf2,xf3,xf4};
			}
		}
	}

}

